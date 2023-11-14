package llist

import "fmt"

// Initialize a new list. The number of elements size.
func New(size int) *LinkedList {

	if size <= 0 {
		fmt.Println("ERROR (in New): list size is set incorrectly")
		return nil
	}

	head := node{}
	current := &head
	for i := 1; i < size; i++ {
		current.next = &node{}
		current = current.next
	}
	return &LinkedList{head: &head}
}
