package dllist

import "fmt"

func (l *DoublyLinkedList) UpdateAt(pos int, val int) {
	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in UpdateAt): list index out of range")
		return
	}
	current := l.head
	for i := 0; i != pos; i++ {
		current = current.next
	}
	current.val = val
}
