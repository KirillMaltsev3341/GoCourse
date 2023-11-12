package bst

func (b *Bst) Size() int {
	size := 0
	b.root.sizeByNode(&size)
	return size
}

func (n *node) sizeByNode(size *int) {
	if n == nil {
		return
	}
	*size++
	n.left.sizeByNode(size)
	n.right.sizeByNode(size)
}
