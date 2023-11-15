package queue

func (q *Queue) Pop() {
	if q.Size() == 1 {
		q.front = nil
		q.back = nil
		return
	}
	if !(q.Empty()) {
		q.front.prev.next = nil
		q.front = q.front.prev
	}
}
