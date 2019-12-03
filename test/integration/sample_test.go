package integration

import (
	"context"
	"strconv"
	"testing"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"
	"github.com/stretchr/testify/assert"
)

func Test_Sample(t *testing.T) {
	for i := 0; i < 100; i++ {
		data := strconv.Itoa(i)
		r, e := client.SampleAPI(context.Background(), &sample.SampleReq{Mess: data})
		assert.Nil(t, e)
		assert.Equal(t, r.Data.Data, data)
	}
}
