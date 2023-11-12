package bst

type node struct {
	key   int
	left  *node
	right *node
}

type Bst struct {
	root *node
}
