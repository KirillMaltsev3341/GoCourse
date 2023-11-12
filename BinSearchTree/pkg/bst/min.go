package bst

func (b *Bst) Min() int {
	return b.root.minByNode()
}

func (n *node) minByNode() int {
	if n.left == nil {
		return n.key
	}
	return n.left.minByNode()
}
