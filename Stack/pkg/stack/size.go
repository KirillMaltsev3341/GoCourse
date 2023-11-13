package stack

func (s *Stack) Size() int {
	size := 0
	tmp := s.top
	for tmp != nil {
		tmp = tmp.next
		size++
	}
	return size
}
