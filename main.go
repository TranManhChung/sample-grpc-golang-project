package main

import (
	"context"
	"log"
	"net"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":56789"
)

type server struct{}

// SampleAPI ...
func (s *server) SampleAPI(ctx context.Context, req *sample.SampleReq) (*sample.SampleRes, error) {

	return &sample.SampleRes{Data: &sample.SampleRes_Data{Data: req.Mess}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
