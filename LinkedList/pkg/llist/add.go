package llist

// Add an item to the end of the list.
func (l *LinkedList) Add(val int) {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &node{val: val, next: nil}
}
