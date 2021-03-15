package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "grpc/sample/config"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewConfigStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &pb.ConfigRequest{Profile: "dev"})
	if err != nil {
		log.Fatalf("could not request: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Config: %v\n", r)
}
