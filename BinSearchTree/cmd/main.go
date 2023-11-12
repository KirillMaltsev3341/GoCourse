package main

import (
	"fmt"

	"main.go/pkg/bst"
)

func main() {
	tree := bst.New(100)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(40)
	tree.Insert(45)
	tree.Insert(35)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(250)
	tree.Insert(150)

	fmt.Println(tree.Min())
	fmt.Println(tree.Max())

	fmt.Println(tree.Size())
	tree.PreorderTraversal()

	fmt.Println(tree.Find(300))
	fmt.Println(tree.Find(500))

	tree.Insert(150)
	tree.Insert(35)
	tree.Remove(90)
	tree.Remove(300)
	tree.Remove(100)
	tree.Remove(10000)

	fmt.Println(tree.Size())
	tree.PreorderTraversal()
	fmt.Println(tree.Find(300))

	tree.Insert(600)

	fmt.Println(tree.Size())
	tree.PreorderTraversal()
}
