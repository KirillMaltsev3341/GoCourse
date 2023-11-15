package stack

func (s *Stack) Push(v int) {
	new := &obj{value: v, next: s.top}
	s.top = new
}
