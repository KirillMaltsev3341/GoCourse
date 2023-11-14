package llist

import "fmt"

// List output
func (l *LinkedList) PrintList() {
	current := l.head

	for current.next != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}
	fmt.Printf("%d\n", current.val)
}
