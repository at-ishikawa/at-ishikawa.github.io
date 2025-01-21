package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runMain() error {
	addrs := []string{
		"redis-node-1:6379",
		"redis-node-2:6379",
		"redis-node-3:6379",
		"redis-node-4:6379",
		"redis-node-5:6379",
		"redis-node-6:6379",
	}
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		// Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Addrs: addrs,
	})
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("rdb.Ping > %w", err)
	}

	// var rdb *redis.Client
	// for _, addr := range addrs {
	// 	rdb = redis.NewClient(&redis.Options{
	// 		Addr: addr,
	// 	})
	// 	defer rdb.Close()
	// 	if err := rdb.Ping(context.Background()).Err(); err != nil {
	// 		return fmt.Errorf("rdb.Ping > %w", err)
	// 	}
	// }

	ctx = context.Background()
	users := []string{"user1", "user2", "user3", "user4", "user5"}
	for _, user := range users {
		redisKey := fmt.Sprintf("user:%s", user)
		for i := 0; i < 5; i++ {
			old, err := rdb.Get(ctx, redisKey).Result()
			if err != nil && err != redis.Nil {
				return fmt.Errorf("rdb.Get > %w", err)
			}

			var val int
			if old != "" {
				val, err = strconv.Atoi(old)
				if err != nil {
					return fmt.Errorf("strconv.Atoi > %w", err)
				}
				val++
			}
			if err := rdb.Set(ctx, redisKey, strconv.Itoa(val), time.Hour).Err(); err != nil {
				return fmt.Errorf("rdb.Set > %w", err)
			}
		}
	}

	return nil
}
