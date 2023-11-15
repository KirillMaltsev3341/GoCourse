package main

import (
	"fmt"

	"main.go/pkg/llist"
)

func main() {
	List := llist.New(4)
	fmt.Println(List.Size())
	List.PrintList()
	fmt.Println()

	List = llist.NewFromSlice([]int{-11, 22, 33, 44, -55, 66, -77})
	fmt.Println(List.Size())
	List.PrintList()
	fmt.Println("[0]:", List.At(0))
	fmt.Println("[3]:", List.At(3))
	fmt.Println("[6]:", List.At(6))
	fmt.Println()

	List.Add(100)
	List.Add(200)
	List.UpdateAt(0, 50)
	List.UpdateAt(1, -243)
	fmt.Println(List.Size())
	List.PrintList()
	fmt.Println()

	List.Pop()
	List.DeleteAt(0)
	List.DeleteAt(3)
	fmt.Println(List.Size())
	List.PrintList()
	fmt.Println()

	x := llist.New(5)
	x.PrintList()
}
