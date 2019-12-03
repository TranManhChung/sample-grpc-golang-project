package integration

import (
	"os"
	"testing"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"

	"google.golang.org/grpc"
)

var client sample.SampleServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:56789", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = sample.NewSampleServiceClient(conn)
	os.Exit(m.Run())
}
