package deque

func (d *Deque) PushFront(val int) {
	new := &obj{value: val, prev: d.front}
	if d.Empty() {
		d.back = new
	} else {
		d.front.next = new
	}
	d.front = new
}
