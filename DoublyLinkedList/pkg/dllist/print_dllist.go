package dllist

import "fmt"

func (l *DoublyLinkedList) PrintDoublyLinkedList() {
	current := l.head
	for current.next != nil {
		fmt.Printf("%d <-> ", current.val)
		current = current.next
	}
	fmt.Printf("%d\n", current.val)
}
