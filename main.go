package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":56789"
)

var queueJob chan string

type server struct{}

// SampleAPI ...
func (s *server) SampleAPI(ctx context.Context, req *sample.SampleReq) (*sample.SampleRes, error) {
	time.Sleep(1 * time.Second)
	queueJob <- req.Mess
	return &sample.SampleRes{Data: &sample.SampleRes_Data{Data: req.Mess}}, nil
}

func main() {
	queueJob = make(chan string, 10)
	for i := 0; i < 8; i++ {
		go func() {
			for j := range queueJob {
				fmt.Println(j)
			}
		}()
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &server{})

	reflection.Register(s)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		fmt.Println("Server now listening")
	}()
	go func() {
		<-sigs
		done <- true
	}()
	fmt.Println("Ctrl-C to interrupt...")
	<-done
}
