package user

import (
	"io/ioutil"
	"log"
	"testing"
)

var dsn = "root:password@/test?charset=utf8&parseTime=true"

func BenchmarkPaginateUsers(b *testing.B) {
	address := ":50052"
	server, err := RunServer(address, dsn, log.New(ioutil.Discard, "", log.LstdFlags))
	if err != nil {
		b.Fatal(err)
	}
	defer server.GracefulStop()
	client := NewGRPCClient(address)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var token string
		for {
			response, err := PaginateUsers(client, "ff", token)
			if err != nil {
				b.Error(err)
			}
			if response != nil {
				token = response.NextToken
			}

			if token == "" {
				break
			}
		}
	}
}

func BenchmarkStreamUsers(b *testing.B) {
	benchCases := []struct {
		name        string
		concurrency uint32
	}{
		{
			name:        "concurrency = 1",
			concurrency: 1,
		},
		{
			name:        "concurrency = 10",
			concurrency: 10,
		},
		{
			name:        "concurrency = 100",
			concurrency: 100,
		},
	}
	address := ":50053"
	server, err := RunServer(address, dsn, log.New(ioutil.Discard, "", log.LstdFlags))
	if err != nil {
		b.Fatal(err)
	}
	defer server.GracefulStop()
	client := NewGRPCClient(address)
	b.ResetTimer()

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := StreamUsers(client, "ff", bc.concurrency)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}

}
