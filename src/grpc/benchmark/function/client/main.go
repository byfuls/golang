package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "byfuls.com/generate/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewProtoParamClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.Get(ctx, &pb.ConfigRequest{Profile: "dev"})
	r, err := c.PostStressUp(ctx, &pb.ProtoRequest{Pid: "pid", GateId: "gateId"})
	if err != nil {
		log.Fatalf("could not request: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Config: %v\n", r)
}
