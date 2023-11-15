package dllist

import "fmt"

func (l *DoublyLinkedList) InsertBefore(pos int, val int) {
	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in InsertBefore): list index out of range")
		return
	}
	if pos == 0 {
		l.PushFront(val)
		return
	}
	current := l.head
	for i := 0; i != pos; i++ {
		current = current.next
	}
	newNode := &node{val: val, next: current, prev: current.prev}
	current.prev.next = newNode
	current.prev = newNode
}
