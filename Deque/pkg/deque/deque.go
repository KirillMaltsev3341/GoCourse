package deque

type obj struct {
	value int
	next  *obj
	prev  *obj
}

// FIFO
type Deque struct {
	front *obj
	back  *obj
}
