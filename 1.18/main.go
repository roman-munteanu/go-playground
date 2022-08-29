package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("-----------------")
	var a, b, c car
	a = "Ferrari"
	b = "Alfa Romeo"
	c = "Maserati"
	Ride([]Vehicle{a, b, c})

	fmt.Println("-----------------")
	sl := []int{1, 2, 3, 4}
	mapped := mapElements(sl, func(n int) int {
		return n * 2
	})
	fmt.Println(mapped)

	fmt.Println("-----------------")
	filtered := filterElements(sl, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(filtered)

	fmt.Println("-----------------")
	forEach(sl, func(el int) {
		fmt.Println(el)
	})
}

// any type
func printAnything[T any](v T) {
	fmt.Println(v)
}

// print the elements of any slice
func printElements[T any](sl []T) {
	for _, v := range sl {
		fmt.Print(v)
	}
}

// explicit types
func increase[T int64 | float64](v T) {
	val := int(v) + 1
	fmt.Println(val)
}

// type constaints
type MyNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumNumbers[T MyNumber](numbers []T) T {
	var result T
	for _, n := range numbers {
		result += n
	}

	return result
}

func max[T MyNumber](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// comparable - Equal("abc", "abc")
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Type fields in structs
type MyEntity[T any] struct {
	inner T
}

func (e *MyEntity[T]) Get() T {
	return e.inner
}

func (e *MyEntity[T]) Set(v T) {
	e.inner = v
}

// generic types
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() {
	n := len(s.items)
	if n <= 0 {
		log.Fatalln("Empty stack")
	}
	s.items = s.items[:n-1]
}

// maps
func extractKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Invoke custom methods using Generics
type Vehicle interface {
	Drive()
}

type car string

func (c car) Drive() {
	fmt.Printf("Driving a %s\n", c)
}

func Ride[T Vehicle](autos []T) {
	for _, v := range autos {
		v.Drive()
	}
}

// alternative:
// func RideAlt(autos []Vehicle) {
// 	for _, v := range autos {
// 		v.Drive()
// 	}
// }

// map elements of a slice
func mapElements[T any](sl []T, f func(el T) T) []T {
	res := make([]T, len(sl))
	for i, v := range sl {
		res[i] = f(v)
	}
	return res
}

func apply[T any](sl []T, f func(el T) T) {
	for i, v := range sl {
		sl[i] = f(v)
	}
}

// filter
func filterElements[T any](sl []T, predicate func(el T) bool) []T {
	res := make([]T, 0)
	for _, v := range sl {
		if predicate(v) {
			res = append(res, v)
		}
	}

	return res
}

// forEach
func forEach[T any](sl []T, consumer func(el T)) {
	for _, v := range sl {
		consumer(v)
	}
}
