package llist

import "fmt"

// Getting an item from the pos position.
func (l *LinkedList) At(pos int) int {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in At): list index out of range")
		return 0
	}

	current := l.head
	for i := 0; i != pos && current.next != nil; i++ {
		current = current.next
	}

	return current.val
}
