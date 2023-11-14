package llist

import "fmt"

// Remove an item from the pos position.
func (l *LinkedList) DeleteAt(pos int) {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in DeleteAt): list index out of range")
		return
	}

	if l.head.next == nil {
		fmt.Println("ERROR (in DeleteAt): Unable to delete element, list size is already 1")
		return
	}

	if pos == 0 {
		l.head = l.head.next
		return
	}

	current := l.head
	for i := 0; i != pos-1 && current.next != nil; i++ {
		current = current.next
	}

	current.next = current.next.next
}
