package main

import (
	"fmt"
	"time"
)

func main() {
	buffCh := make(chan int, 2)
	buffCh <- 1 // will be shown first
	buffCh <- 2
	// buffCh <- 3 // deadlock - capacity 2
	fmt.Println(<-buffCh)
	fmt.Println(<-buffCh)
	// fmt.Println(<-buffCh) // deadlock (blocks when full or empty)

	select {
	case a := <-waitAndSend(10, 2):
		fmt.Println(a)
	case b := <-waitAndSend(20, 1):
		fmt.Println(b)
		// default:
		// 	fmt.Println("This is executed faster - default when channels are too slow")
	}

}

func waitAndSend(x, n int) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(n) * time.Second)
		ch <- x
	}()

	return ch
}
