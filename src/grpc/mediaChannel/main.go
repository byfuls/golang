package main

import (
	"context"
	_ "io"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "grpc/mediaChannel/config"
)

const (
	serverAddress = "127.0.0.1:50051"
)

type server struct {
	pb.UnimplementedMediaChannelLauncherServer
}

func (s *server) LaunchMediaChannel(ctx context.Context, in *pb.MediaChannelRequest) (*pb.MediaChannelResponse, error) {
	log.Printf("[server] received: %v\n", in.GetCommand())

	return &pb.MediaChannelResponse{Address: "get command: " + in.GetCommand()}, nil
}

func (s *server) StreamMediaChannel(in *pb.MediaChannelStreamRequestMessage, srv pb.MediaChannelLauncher_StreamMediaChannelServer) error {

	log.Println("[server] stream response for id: ", in.Id)

	for {
		time.Sleep(5 * time.Second)
		resp := pb.MediaChannelStreamResponseMessage{
			Id:      "test",
			Type:    "type",
			Message: "message",
		}
		if err := srv.Send(&resp); err != nil {
			log.Println("[server] send error: %v\n", err)
			return err
		}

		log.Println("[server] ing...")
	}

	return nil
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
