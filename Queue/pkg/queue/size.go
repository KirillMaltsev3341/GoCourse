package queue

func (q *Queue) Size() int {
	size := 0
	tmp := q.back
	for tmp != nil {
		tmp = tmp.next
		size++
	}
	return size
}
