package bst

func (b *Bst) Remove(k int) {
	b.root = b.root.removeByNode(k)
}

func (n *node) removeByNode(k int) *node {
	if n == nil {
		return n
	}
	if n.key < k {
		n.right = n.right.removeByNode(k)
	} else if n.key > k {
		n.left = n.left.removeByNode(k)
	} else {
		if n.left == nil {
			return n.right
		} else {
			n.key = n.left.maxByNode()
			n.left = n.left.removeByNode(n.key)
		}
	}
	return n
}
