package main

import (
	"fmt"

	"main.go/pkg/queue"
)

func main() {
	q := queue.Queue{}
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(40)
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

	q.Pop()
	q.Pop()
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

	q.Push(50)
	q.Push(60)
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

	q.Pop()
	fmt.Println(q.Size())
	fmt.Println(q.Front(), q.Back())

}
