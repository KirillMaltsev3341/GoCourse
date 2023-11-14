package llist

import "fmt"

// Make the value at the pos position equal to val.
func (l *LinkedList) UpdateAt(pos int, val int) {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in UpdateAt): list index out of range")
		return
	}

	current := l.head
	for i := 0; i != pos && current.next != nil; i++ {
		current = current.next
	}

	current.val = val
}
