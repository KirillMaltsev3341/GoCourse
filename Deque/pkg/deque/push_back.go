package deque

func (d *Deque) PushBack(val int) {
	new := &obj{value: val, next: d.back}
	if d.Empty() {
		d.front = new
	} else {
		d.back.prev = new
	}
	d.back = new
}
