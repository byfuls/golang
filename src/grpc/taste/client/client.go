package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "grpc/taste/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewGRPCFuncsListClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var x int32
	var y int32
	x = 5
	y = 3
	res, err := client.Sum(ctx, &pb.SumArgs{Value1: x, Value2: y})
	if err != nil {
		log.Fatalf("could not request: %v\n", err)
	}

	log.Printf("response: %v\n", res)
}
