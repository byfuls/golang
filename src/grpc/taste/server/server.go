package main

import (
	"context"
	pb "grpc/taste/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGRPCFuncsListServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumArgs) (*pb.SumReturns, error) {
	log.Println(in)
	calc := in.Value1 + in.Value2
	return &pb.SumReturns{Value: calc}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterGRPCFuncsListServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
