package dllist

import "fmt"

func (l *DoublyLinkedList) PopBack() {
	if l.head.next == nil {
		fmt.Println("ERROR (in PopBack): Unable to delete element, list size is already 1")
		return
	}
	l.tail = l.tail.prev
	l.tail.next = nil
}
