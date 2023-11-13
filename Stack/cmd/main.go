package main

import (
	"fmt"

	"main.go/pkg/stack"
)

func main() {
	s := stack.Stack{}
	fmt.Println(s)

	s.Push(10)
	s.Push(20)

	s.Pop()
	s.Pop()
	s.Pop()

	fmt.Println(s.Top())
	fmt.Println(s.Empty())
	fmt.Println(s.Size())

	fmt.Println("----------")
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Top())
	fmt.Println(s.Empty())
	fmt.Println(s.Size())

	s.Pop()
	fmt.Println("----------")
	fmt.Println(s.Top())
	fmt.Println(s.Empty())
	fmt.Println(s.Size())
}
