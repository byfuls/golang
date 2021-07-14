package main

import (
	"context"
	"log"
	"testing"

	pb "byfuls.com/generate/message"

	"google.golang.org/grpc"
)

func TestMain(t *testing.T) {
	conn, err := grpc.Dial(":2219", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v\n", err)
	}
	defer conn.Close()

	client := pb.NewMessageStreamClient(conn)
	stream, err := client.StreamWriter(context.Background())
	if err != nil {
		log.Fatalf("can not receive %v\n", err)
	}
	// defer func() {
	// 	_, err := stream.CloseAndRecv()
	// 	if err != nil {
	// 		log.Fatalf("stream close and recv error %v\n", err)
	// 	}
	// }()

	for i := 0; i < 5; i++ {
		err = stream.Send(&pb.ReqPkt{Command: "Hello, World!"})
		if err != nil {
			log.Fatalf("stream send error %v\n", err)
		}

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream close and recv error %v\n", err)
	}
	log.Println("close and recv: ", res)
}

func BenchmarkMain(b *testing.B) {
	conn, err := grpc.Dial(":2219", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v\n", err)
	}
	defer conn.Close()

	client := pb.NewMessageStreamClient(conn)
	stream, err := client.StreamWriter(context.Background())
	if err != nil {
		log.Fatalf("can not receive %v\n", err)
	}
	// defer func() {
	// 	_, err := stream.CloseAndRecv()
	// 	if err != nil {
	// 		log.Fatalf("stream close and recv error %v\n", err)
	// 	}
	// }()

	for i := 0; i < b.N; i++ {
		err = stream.Send(&pb.ReqPkt{Command: "Hello, World!"})
		if err != nil {
			log.Fatalf("stream send error %v\n", err)
		}

	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream close and recv error %v\n", err)
	}
	// log.Println("close and recv: ", res)

}
