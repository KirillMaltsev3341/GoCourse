package stack

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}
	s.top = s.top.next
}
