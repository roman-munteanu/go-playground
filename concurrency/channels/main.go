package main

import (
	"fmt"
	"time"
)

func main() {
	quitChannel := make(chan bool)
	go execute(quitChannel)
	fmt.Println("Inside main goroutine")
	v := <-quitChannel
	fmt.Println(v)

	ch := make(chan int)
	go send(ch)
	for x := range ch {
		fmt.Println(x)
	}
	_, ok := <-ch
	fmt.Println(ok) // channel is closed
}

func execute(ch chan bool) {
	fmt.Println("Execute called")
	ch <- true
}

func send(ch chan int) {
	i := 0
	for i < 4 {
		ch <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
