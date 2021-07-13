package main

import (
	"context"
	"log"
	"os"
	"testing"

	pb "byfuls.com/generate/proto"
	"google.golang.org/grpc"
)

func BenchmarkMain(b *testing.B) {
	// conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewProtoParamClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// var tmp int64 = 0
	for i := 0; i < b.N; i++ {
		// defer func() {
		// 	fmt.Println("rslt: ", tmp)
		// }()

		// r, err := c.PostStressUp(ctx, &pb.ProtoRequest{Pid: "pid", GateId: "gateId"})
		_, err := c.PostStressUp(context.Background(), &pb.ProtoRequest{Pid: "pid", GateId: "gateId"})
		if err != nil {
			log.Fatalf("could not request: %v\n", err)
			os.Exit(1)
		}
		// tmp += int64(r.TotalStress)
	}

	// r, err := c.PostStressUp(ctx, &pb.ProtoRequest{Pid: "pid", GateId: "gateId"})
	// if err != nil {
	// 	log.Fatalf("could not request: %v\n", err)
	// 	os.Exit(1)
	// }
	// log.Printf("Config: %v\n", r)
}
