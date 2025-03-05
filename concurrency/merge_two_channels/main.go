package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	result := make(chan int)

	ch1 := make(chan int)
	ch2 := make(chan int)

	channels := []chan int{ch1, ch2}

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	fillUp := func(ch chan int) {
		for v := range ch {
			result <- v
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go fillUp(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println(v)
	}
}
