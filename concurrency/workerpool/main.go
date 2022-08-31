package main

import (
	"context"
	"errors"
	"fmt"
)

type Operation string

const (
	Add      Operation = "+"
	Multiply           = "*"
)

type WorkRequest struct {
	Op   Operation
	Val1 int
	Val2 int
}

type WorkResponse struct {
	Req    *WorkRequest
	Result int
	Err    error
}

func main() {
	numberOfWorkers := 5
	in := make(chan WorkRequest, 5)
	out := make(chan WorkResponse, 5)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < numberOfWorkers; i++ {
		go worker(ctx, 1000+i, in, out)
	}

	for i := 0; i < 5; i++ {
		in <- WorkRequest{Op: Add, Val1: 10, Val2: i}
	}

	for i := 0; i < 5; i++ {
		in <- WorkRequest{Op: Multiply, Val1: 100, Val2: i}
	}

	for i := 0; i < 10; i++ {
		resp := <-out
		if resp.Err != nil {
			panic(resp.Err)
		}
		fmt.Printf("Operation: %s, result: %v\n", resp.Req.Op, resp.Result)
	}
}

func worker(ctx context.Context, id int, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-in:
			fmt.Printf("Worker ID: %d performing work on operation: %v\n", id, req.Op)
			out <- *process(&req)
		}
	}
}

func process(req *WorkRequest) *WorkResponse {
	resp := &WorkResponse{
		Req: req,
	}
	switch req.Op {
	case Add:
		resp.Result = req.Val1 + req.Val2
	case Multiply:
		resp.Result = req.Val1 * req.Val2
	default:
		resp.Err = errors.New("unsupported operation")
	}

	return resp
}
