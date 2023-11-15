package stack

type obj struct {
	value int
	next  *obj
}

// LIFO
type Stack struct {
	top *obj
}
