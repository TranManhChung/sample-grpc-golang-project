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
	port      = ":56789"
	maxWorker = 8
)

var queueJob chan string

type server struct{}

// SampleAPI ...
func (s *server) SampleAPI(ctx context.Context, req *sample.SampleReq) (*sample.SampleRes, error) {
	queueJob <- req.Mess
	return &sample.SampleRes{Data: &sample.SampleRes_Data{Data: req.Mess}}, nil
}

func main() {
	done := make(chan bool, 1)
	var stopSig []chan bool

	queueJob = make(chan string, 100)
	var job string
	for i := 0; i < maxWorker; i++ {
		a := make(chan bool, 1)
		stopSig = append(stopSig, a)
		go func(i int) {
			for {
				time.Sleep(1 * time.Second)
				select {
				case job = <-queueJob:
					fmt.Println(job)
				case stop := <-done:
					if stop == false && len(queueJob) == 0 {
						stopSig[i] <- true
						return
					}
				}
			}
		}(i)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &server{})
	reflection.Register(s)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	go func() {
		<-sigs
		s.Stop()
		close(done)
	}()
	fmt.Println("Ctrl-C to interrupt...")
	for i := 0; i < maxWorker; i++ {
		<-stopSig[i]
	}

}
