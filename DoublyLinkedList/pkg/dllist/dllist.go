package dllist

type node struct {
	val  int
	next *node
	prev *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
}
