package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/jmoiron/sqlx"

	"github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/internal/user"
)

var (
	dsn = flag.String("dsn", "root:password@/test", "DSN")
)

func main() {
	flag.Parse()
	db, err := sqlx.Connect("mysql", *dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ctx := context.Background()

	rand.Seed(time.Now().UnixNano())
	const chunkSize = 1000
	testSize := 1000 * chunkSize
	for i := 0; i < testSize/chunkSize; i++ {
		records := make([]user.Record, chunkSize)
		for j := 0; j < chunkSize; j++ {
			name := petname.Generate(rand.Intn(3)+3, "")
			records[j] = user.Record{
				Name:      name,
				Email:     name + "@gmail.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		}
		_, err := db.NamedExecContext(ctx, "INSERT INTO users (name, email, created_at, updated_at) VALUES (:name, :email, :created_at, :updated_at)", records)
		if err != nil {
			log.Println(err)
		}
		<-time.After(1 * time.Millisecond)
		log.Printf("Current: %d, MaxSize: %d\n", (i+1)*chunkSize, testSize)
	}
}
