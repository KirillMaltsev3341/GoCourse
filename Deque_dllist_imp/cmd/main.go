package main

import (
	"fmt"

	"main.go/pkg/deque"
)

func main() {
	d := deque.Deque{}
	fmt.Println(d.Size())
	fmt.Println(d.Front(), d.Back())
	fmt.Println("--------------------------------------------")

	d.PopBack()
	d.PopFront()
	fmt.Println(d.Size())
	fmt.Println(d.Front(), d.Back())
	fmt.Println("--------------------------------------------")

	d.PushFront(14)
	d.PushFront(29)
	d.PushBack(-8)
	d.PushBack(128)
	d.PushBack(54)
	d.PushFront(453)

	fmt.Println(d.Size())
	fmt.Println(d.Front(), d.Back())
	fmt.Println("--------------------------------------------")

	d.PopBack()
	d.PopBack()
	d.PopFront()
	fmt.Println(d.Size())
	fmt.Println(d.Front(), d.Back())
	fmt.Println("--------------------------------------------")
}
