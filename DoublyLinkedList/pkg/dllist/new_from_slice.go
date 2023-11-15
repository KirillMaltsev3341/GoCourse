package dllist

import "fmt"

func NewFromSlice(s []int) *DoublyLinkedList {
	if len(s) == 0 {
		fmt.Println("ERROR (in NewFromSlise): Unable convert an empty slice to a list")
		return nil
	}

	head := node{val: s[0]}
	current := &head
	for i := 1; i < len(s); i++ {
		current.next = &node{val: s[i], prev: current}
		current = current.next
	}
	return &DoublyLinkedList{head: &head, tail: current}
}
