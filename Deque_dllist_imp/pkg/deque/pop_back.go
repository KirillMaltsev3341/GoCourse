package deque

func (d *Deque) PopBack() {
	if d.Size() == 1 {
		d.front = nil
		d.back = nil
	} else if !(d.Empty()) {
		d.back.next.prev = nil
		d.back = d.back.next
	}
}
