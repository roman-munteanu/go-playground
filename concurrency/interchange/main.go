package main

import (
	"context"
	"errors"
	"fmt"
)

type WorkRequest struct {
	Value int
}

type WorkResponse struct {
	Req    *WorkRequest
	Result int
	Err    error
}

func main() {
	in := make(chan *WorkRequest, 5)
	out := make(chan *WorkResponse, 5)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go doWork(ctx, in, out)

	in <- &WorkRequest{Value: 10}
	in <- &WorkRequest{Value: -10}
	in <- &WorkRequest{Value: 20}

	for i := 0; i < 3; i++ {
		resp := <-out
		fmt.Printf("Request: %v, Result: %v, Error: %v\n", resp.Req, resp.Result, resp.Err)
	}
}

func doWork(ctx context.Context, in chan *WorkRequest, out chan *WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-in:
			out <- process(req)
		}
	}
}

func process(req *WorkRequest) *WorkResponse {
	resp := WorkResponse{
		Req: req,
	}

	if req.Value < 0 {
		resp.Err = errors.New("unsupported value")
		return &resp
	}

	resp.Result = req.Value * 2
	return &resp
}
