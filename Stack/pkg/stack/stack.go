package stack

type obj struct {
	value int
	next  *obj
}

type Stack struct {
	top *obj
}
