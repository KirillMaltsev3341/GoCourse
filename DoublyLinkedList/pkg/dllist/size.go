package dllist

func (l *DoublyLinkedList) Size() int {
	size := 1
	current := l.head
	for current.next != nil {
		current = current.next
		size++
	}
	return size
}
