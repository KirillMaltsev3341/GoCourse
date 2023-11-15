package dllist

import "fmt"

func (l *DoublyLinkedList) PopFront() {
	if l.head.next == nil {
		fmt.Println("ERROR (in PopFront): Unable to delete element, list size is already 1")
		return
	}
	l.head = l.head.next
	l.head.prev = nil
}
