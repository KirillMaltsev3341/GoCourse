package dllist

func (l *DoublyLinkedList) PushFront(val int) {
	l.head.prev = &node{val: val, next: l.head}
	l.head = l.head.prev
}
