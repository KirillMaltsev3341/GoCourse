package llist

import "fmt"

// Creating a linked list from a slice.
func NewFromSlice(s []int) *LinkedList {

	if len(s) == 0 {
		fmt.Println("ERROR (in NewFromSlise): Unable convert an empty slice to a list")
		return nil
	}

	head := node{val: s[0]}
	current := &head
	for i := 0; i < len(s)-1; i++ {
		current.next = &node{val: s[i+1]}
		current = current.next
	}
	return &LinkedList{head: &head}
}
