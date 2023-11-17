package deque

func (d *Deque) PopFront() {
	if d.Size() == 1 {
		d.front = nil
		d.back = nil
	} else if !(d.Empty()) {
		d.front.prev.next = nil
		d.front = d.front.prev
	}
}
