package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go work(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM)

	// main blocked until Ctrl+C
	<-signalCh
}

func task() {
	fmt.Println("task execution")
}

func work(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			ticker.Stop()
			return
		case <-ticker.C:
			task()
		}
	}
}
