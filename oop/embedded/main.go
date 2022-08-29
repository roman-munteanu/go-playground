package main

import "fmt"

type alpha struct {
	field1 int
	field2 int
}

type beta struct {
	field1 int
	field3 int
}

type gamma struct {
	a alpha
	b beta
}

func (a alpha) do() {
	fmt.Println("do for alpha")
}

func (b beta) do() {
	fmt.Println("do for beta")
}

func main() {
	var g gamma
	g.a.do()
	g.b.do()
}
