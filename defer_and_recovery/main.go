package main

import "fmt"

func main() {
	defer fmt.Println("One - invoked second")
	defer fmt.Println("Two - invoked first")
	fmt.Println("Before panic")
	gonnaPanic()
	fmt.Println("After panic")
}

func gonnaPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered")
		}
	}()
	panic("Some error happened") // similar to exceptions
}
