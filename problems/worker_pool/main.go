package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	numberOfWorkers := 3
	numberOfItemsToProcess := 20
	ch := make(chan int)
	doneCh := make(chan any)

	go func() {
		for i := 0; i < numberOfItemsToProcess; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < numberOfWorkers; i++ {
		go worker(ctx, i, ch, doneCh)
	}

	for i := 0; i < numberOfWorkers; i++ {
		<-doneCh
	}

	fmt.Println("All done")
}

func worker(ctx context.Context, workerNum int, ch <-chan int, doneCh chan<- any) {
	for {
		select {
		case val := <-ch:
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("worker: %d - processed value: %d\n", workerNum, val)
		case <-ctx.Done():
			fmt.Printf("worker: %d - done\n", workerNum)
			doneCh <- struct{}{}
			return
		}
	}
}
