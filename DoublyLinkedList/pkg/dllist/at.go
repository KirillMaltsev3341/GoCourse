package dllist

import "fmt"

func (l *DoublyLinkedList) At(pos int) int {
	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in At): list index out of range")
		return 0
	}

	current := l.head
	for i := 0; i != pos; i++ {
		current = current.next
	}
	return current.val
}
