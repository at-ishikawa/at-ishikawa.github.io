package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/protos"
)

const (
	grpcPort = ":50051"
)

type Server struct {
	pb.HelloWorldServer
}

func (s *Server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("receive name: %s in SayHello", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func (s *Server) KeepReplyingHello(in *pb.HelloRequest, stream pb.HelloWorld_KeepReplyingHelloServer) error {
	startTime := time.Now()
	if err := stream.SetHeader(metadata.New(map[string]string{
		"start": startTime.Format(time.RFC3339),
	})); err != nil {
		log.Printf("err to send header: %+v", err)
		return status.Errorf(codes.Internal, "error to set header: %+v", err)
	}

	for i := 1; i < 10; i++ {
		err := stream.Send(&pb.HelloResponse{
			Message: fmt.Sprintf("%d: Hello %s", i, in.Name),
		})
		if err != nil {
			log.Printf("error: %+v", err)
			return status.Errorf(codes.Internal, "error to Send: %+v", err)
		}
	}
	duration := time.Now().Sub(startTime)
	stream.SetTrailer(metadata.New(map[string]string{
		"duration": duration.String(),
	}))
	log.Printf("receive name: %s in KeepReplyingHello, took: %d", in.Name, duration)
	return nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
