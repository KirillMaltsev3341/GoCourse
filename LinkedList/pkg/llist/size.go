package llist

// Return the length of the list
func (l *LinkedList) Size() int {
	size := 1
	current := l.head
	for current.next != nil {
		size++
		current = current.next
	}
	return size
}
