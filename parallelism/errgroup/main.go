package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

var errTest = errors.New("test error")

func main() {
	group := errgroup.Group{}

	group.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("first task")
		return nil
	})

	group.Go(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("second task")
		return nil
	})

	group.Go(func() error {
		fmt.Println("third task")
		return errTest
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("error occured: %v\n", err)
		return
	}

	fmt.Println("all tasks done")
}
