package main

import (
	"context"
	pb "grpc/sample/config"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConfigStoreServer
}

func (s *server) Get(ctx context.Context, in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	log.Printf("Received profile: %v\n", in.GetProfile())
	return &pb.ConfigResponse{JsonConfig: `"{"main":"http://google.com"}"`}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterConfigStoreServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		os.Exit(1)
	}
}
