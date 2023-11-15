package dllist

func (l *DoublyLinkedList) PushBack(val int) {
	l.tail.next = &node{val: val, prev: l.tail}
	l.tail = l.tail.next
}
