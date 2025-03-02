package main

import "fmt"

type Stack struct {
	ls []int
}

func (s *Stack) Push(val int) {
	s.ls = append(s.ls, val)
}

func (s *Stack) Pop() int {
	idx := len(s.ls) - 1
	val := s.ls[idx]
	s.ls = s.ls[:idx]
	return val
}

func (s Stack) IsEmpty() bool {
	return len(s.ls) == 0
}

func (s Stack) String() string {
	return fmt.Sprintf("%v\n", s.ls)
}

func main() {
	st := Stack{}

	st.Push(10)
	st.Push(20)
	st.Push(30)

	fmt.Println(st)

	fmt.Println(st.Pop())
	fmt.Println(st.Pop())

	fmt.Println(st.IsEmpty())
	fmt.Println(st)
}
