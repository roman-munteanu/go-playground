package main

import (
	"fmt"
)

type Queue struct {
	ls []int
}

func (q *Queue) Enqueue(val int) {
	q.ls = append(q.ls, val)
}

func (q *Queue) Dequeue() int {
	val := q.ls[0]
	q.ls = q.ls[1:len(q.ls)]
	return val
}

func (q Queue) IsEmpty() bool {
	return len(q.ls) == 0
}

func (q Queue) String() string {
	return fmt.Sprintf("%v\n", q.ls)
}

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q)
}
