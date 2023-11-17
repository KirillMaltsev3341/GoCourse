package deque

func (d *Deque) Size() int {
	tmp := d.back
	size := 0
	for tmp != nil {
		tmp = tmp.next
		size++
	}
	return size
}
