package dllist

import "fmt"

func (l *DoublyLinkedList) InsertAfter(pos int, val int) {
	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in InsertAfter): list index out of range")
		return
	}
	if pos == l.Size()-1 {
		l.PushBack(val)
		return
	}
	current := l.head
	for i := 0; i != pos; i++ {
		current = current.next
	}
	newNode := &node{val: val, next: current.next, prev: current}
	current.next.prev = newNode
	current.next = newNode
}
