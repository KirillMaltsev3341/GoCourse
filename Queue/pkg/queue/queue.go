package queue

type obj struct {
	value int
	next  *obj
	prev  *obj
}

// FIFO
type Queue struct {
	front *obj
	back  *obj
}
