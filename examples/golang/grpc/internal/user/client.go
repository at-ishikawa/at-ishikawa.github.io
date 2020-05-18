package user

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"

	pb "github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/protos"
)

func PaginateUsers(client pb.UserServiceClient, keyword string, token string) (*pb.PaginateUsersResponse, error) {
	response, err := client.PaginateUsers(context.Background(), &pb.PaginateUsersRequest{
		Keyword:   keyword,
		PageToken: token,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func StreamUsers(client pb.UserServiceClient, keyword string, concurrency uint32) ([]*pb.StreamUsersResponse, error) {
	stream, err := client.StreamUsers(context.Background(), &pb.StreamUsersRequest{
		Keyword:     keyword,
		Concurrency: concurrency,
	})
	if err != nil {
		return nil, err
	}
	defer stream.CloseSend()

	var responses []*pb.StreamUsersResponse
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("failed to receive some response: %+v\n", err)
			return responses, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func NewGRPCClient(target string) pb.UserServiceClient {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewUserServiceClient(conn)
	return client
}
