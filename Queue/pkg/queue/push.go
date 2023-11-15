package queue

func (q *Queue) Push(val int) {
	new := &obj{value: val, next: q.back}
	if q.Empty() {
		q.front = new
	} else {
		q.back.prev = new
	}
	q.back = new

}
