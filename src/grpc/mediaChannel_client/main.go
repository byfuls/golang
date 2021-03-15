package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "grpc/mediaChannel/config"
)

const (
	serverAddress = "127.0.0.1:50051"

	defaultCommand = "start"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[client] did not connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewMediaChannelLauncherClient(conn)

	command := defaultCommand
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.LaunchMedialChannel(ctx, &pb.MediaChannelRequest{Command: command})
	if err != nil {
		log.Fatalf("[client] could not launch media channel: %v\n", err)
		os.Exit(1)
	}

	log.Printf("[client] get response: %v\n", r.GetAddress())
}
