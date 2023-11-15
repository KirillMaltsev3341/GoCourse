package queue

func (q *Queue) Empty() bool {
	return q.back == nil
}
