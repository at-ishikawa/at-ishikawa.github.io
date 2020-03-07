package main

import (
	"fmt"
	"io"
	"os"
	"time"

	pb "github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	serverAddr = "127.0.0.1"
	grpcPort   = ":50051"
)

func SayHello(client pb.HelloWorldClient) error {
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "client name",
	})
	if err != nil {
		return err
	}
	fmt.Printf("response: %v\n", response)
	return nil
}

func KeepGettingHello(client pb.HelloWorldClient) error {
	startTime := time.Now()
	stream, err := client.KeepReplyingHello(context.Background(), &pb.HelloRequest{
		Name: "client name",
	})
	if err != nil {
		return err
	}
	md, err := stream.Header()
	if err != nil {
		return err
	}
	fmt.Printf("header: %+v\n", md)

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("failed to receive some response: %+v\n", err)
			return stream.CloseSend()
		}
		fmt.Printf("response: %v\n", response)
		// <-time.After(100 * time.Millisecond)
	}
	md = stream.Trailer()
	duration := time.Now().Sub(startTime)
	fmt.Printf("trailer: %+v\n", md)
	fmt.Printf("client duration: %+v\n", duration)
	return nil
}

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(serverAddr+grpcPort, opts...)
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)
	if err := SayHello(client); err != nil {
		fmt.Printf("main > SayHello error: %+v\n", err)
		os.Exit(1)
	}
	if err := KeepGettingHello(client); err != nil {
		fmt.Printf("main > KeepGettingHello error: %+v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
