package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "grpc/mediaChannel/config"
)

const (
	serverAddress = "127.0.0.1:50051"
)

type server struct {
	pb.UnimplementedMediaChannelLauncherServer
}

func (s *server) LaunchMedialChannel(ctx context.Context, in *pb.MediaChannelRequest) (*pb.MediaChannelResponse, error) {
	log.Printf("[server] received: %v\n", in.GetCommand())

	return &pb.MediaChannelResponse{Address: "get command: " + in.GetCommand()}, nil
}

func main() {
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("[server] failed to listen: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterMediaChannelLauncherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[server] failed to serve: %v\n", err)
		os.Exit(1)
	}
}
