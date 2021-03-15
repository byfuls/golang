package main

import (
	"context"
	"io"
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
	r, err := c.LaunchMediaChannel(ctx, &pb.MediaChannelRequest{Command: command})
	if err != nil {
		log.Fatalf("[client] could not launch media channel: %v\n", err)
		os.Exit(1)
	}

	log.Printf("[client] get response: %v\n", r.GetAddress())

	/****/

	in := &pb.MediaChannelStreamRequestMessage{Id: "hi"}
	stream, err := c.StreamMediaChannel(context.Background(), in)
	if err != nil {
		log.Fatalf("[client] open stream error: %v\n", err)
		os.Exit(1)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true // means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("[client] cannot receive: %v\n", resp.Message)
				os.Exit(1)
			}
			log.Printf("[client] received: %v\n", resp)
		}
	}()

	<-done // we will wait until all response is received
	log.Printf("[client] finished")
}
