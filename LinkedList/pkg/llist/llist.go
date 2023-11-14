package llist

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}
