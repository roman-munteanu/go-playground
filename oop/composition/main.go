package main

import "fmt"

type alpha struct{}

func (a alpha) shared() {
	fmt.Println("shared from alpha")
}

func (a alpha) Execute() {
	a.shared()
}

type beta struct {
	a alpha
}

func (b beta) shared() {
	fmt.Println("shared from beta")
}

func main() {
	alpha{}.Execute()
	beta{}.shared()
	b := beta{}
	c := b.a
	c.Execute()
}
