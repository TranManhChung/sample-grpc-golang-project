package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func Test_Sample(t *testing.T) {
	data := "hello"
	r, e := client.SampleAPI(context.Background(), &sample.SampleReq{Mess: data})
	assert.Nil(t, e)
	assert.Equal(t, r.Data.Data, data)
}

func Test_ParseStatus(t *testing.T) {
	data := ""
	_, e := client.SampleAPI(context.Background(), &sample.SampleReq{Mess: data})
	stt, _ := status.FromError(e)

	out, _ := json.Marshal(e)
	fmt.Println(string(out))

	fmt.Println("status_code : ", stt.Code())
	fmt.Println("status_message : ", stt.Message())
	fmt.Println("status_detail : ", stt.Details())
	fmt.Println("status_data : ", stt.Proto())

	for i, detail := range stt.Details() {
		fmt.Println("detail : ", i, " / ", detail.(*errdetails.ErrorInfo))
	}
}
