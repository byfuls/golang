package main

import (
	"io"
	"log"
	"net"
	"testing"

	pb "byfuls.com/generate/message"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) StreamWriter(stream pb.MessageStream_StreamWriterServer) error {
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer func() {
	// 			wg.Done()
	// 			stream.SendAndClose(&pb.ResPkt{Result: "Done"})
	// 		}()

	for {
		//reqPkt, err := stream.Recv()
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream recv error %v\n", err)
		}
		//log.Println(reqPkt)
	}
	return stream.SendAndClose(&pb.ResPkt{Result: "Done"})

	// 	}()
	// }
	// wg.Wait()
	// return errors.New("StreamWriter Done")
}

func TestMain(t *testing.T) {
	lis, err := net.Listen("tcp", ":2219")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterMessageStreamServer(srv, &server{})
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
