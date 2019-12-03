package integration

import (
	"context"
	"strconv"
	"sync"
	"testing"

	sample "github.com/TranManhChung/sample-grpc-golang-project/grpc-gen"
	"github.com/stretchr/testify/assert"
)

func Test_Sample(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			data := strconv.Itoa(i)
			r, e := client.SampleAPI(context.Background(), &sample.SampleReq{Mess: data})
			assert.Nil(t, e)
			assert.Equal(t, r.Data.Data, data)
		}(i)
	}
	wg.Wait()
}
