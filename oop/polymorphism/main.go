package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type rectangle struct {
	a float64
	b float64
}

type circle struct {
	r float64
}

func (r rectangle) Area() float64 {
	return r.a * r.b
}

func (c circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func calculate(s shape) {
	switch v := s.(type) {
	case rectangle:
		fmt.Println("This is a rectangle!")
	case circle:
		fmt.Printf("%v is a circle!\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
		return
	}

	fmt.Println(fmt.Sprintf("Area: %v", s.Area()))
}

func main() {
	rect := rectangle{a: 2, b: 3}
	calculate(rect)
	circ := circle{r: 2}
	calculate(circ)
}
