package bst

func (b *Bst) Max() int {
	return b.root.maxByNode()
}

func (n *node) maxByNode() int {
	if n.right == nil {
		return n.key
	}
	return n.right.maxByNode()
}
