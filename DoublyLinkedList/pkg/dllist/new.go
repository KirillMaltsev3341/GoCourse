package dllist

func New(size int) *DoublyLinkedList {
	if size <= 0 {
		return nil
	}

	head := node{}
	current := &head
	for i := 1; i < size; i++ {
		current.next = &node{prev: current}
		current = current.next
	}
	return &DoublyLinkedList{head: &head, tail: current}
}
