package main

import (
	"fmt"

	"main.go/pkg/dllist"
)

func main() {
	l := dllist.New(5)

	l.PrintDoublyLinkedList()
	fmt.Println(l.Size())

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l = dllist.NewFromSlice(a)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.UpdateAt(3, 100)
	l.PrintDoublyLinkedList()

	l.PopBack()
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.PopFront()
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.PushBack(-25)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.PushFront(-555)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.RemoveAt(4)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.InsertAfter(3, 1111)
	l.InsertAfter(9, 1111)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()

	l.InsertBefore(3, 2222)
	l.InsertBefore(10, 3333)
	l.InsertBefore(12, 2222)
	l.InsertBefore(0, 2222)
	fmt.Println(l.Size())
	l.PrintDoublyLinkedList()
}
