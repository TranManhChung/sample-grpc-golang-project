package main

import (
	"context"
	"log"
	"net"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":56789"
)

type server struct{}

// SampleAPI ...
func (s *server) SampleAPI(ctx context.Context, req *sample.SampleReq) (*sample.SampleRes, error) {
	if req.GetMess() == "" {
		errInfo := &errdetails.ErrorInfo{
			Type:   "INVALID_PARAM",
			Domain: "sample.grpc",
			Metadata: map[string]string{
				"h5": "h5 token is empty",
			},
		}
		clientMsg := &errdetails.LocalizedMessage{
			Locale:  "vietnamese",
			Message: "Thông tin không hợp lệ",
		}
		stt, _ := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).WithDetails(
			errInfo,
			clientMsg,
		)
		return nil, stt.Err()
	}
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
