package bst

func (b *Bst) Find(k int) bool {
	return b.root.findByNode(k)
}

func (n *node) findByNode(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		return n.right.findByNode(k)
	} else if n.key > k {
		return n.left.findByNode(k)
	}
	return true
}
