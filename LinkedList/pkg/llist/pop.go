package llist

import "fmt"

// Remove an element from the end.
func (l *LinkedList) Pop() {

	if l.head.next == nil {
		fmt.Println("ERROR (in Pop): Unable to delete element, list size is already 1")
		return
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}
