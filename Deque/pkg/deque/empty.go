package deque

func (d *Deque) Empty() bool {
	return d.back == nil
}
