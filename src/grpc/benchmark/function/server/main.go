package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "byfuls.com/generate/proto"

	"google.golang.org/grpc"
)

var gVal int64 = 0

type server struct{}

func (s *server) PostStressUp(ctx context.Context, in *pb.ProtoRequest) (*pb.ProtoResponse, error) {
	gVal += 1
	return &pb.ProtoResponse{TotalStress: gVal}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterProtoParamServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		os.Exit(1)
	}
}
