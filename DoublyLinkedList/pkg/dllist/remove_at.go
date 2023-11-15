package dllist

import "fmt"

func (l *DoublyLinkedList) RemoveAt(pos int) {
	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in RemoveAt): list index out of range")
		return
	}
	if l.head.next == nil {
		fmt.Println("ERROR (in RemoveAt): Unable to delete element, list size is already 1")
		return
	}
	if pos == 0 {
		l.PopFront()
		return
	}
	if pos == l.Size()-1 {
		l.PopBack()
		return
	}

	current := l.head
	for i := 0; i != pos-1; i++ {
		current = current.next
	}
	current.next.next.prev = current
	current.next = current.next.next
}
