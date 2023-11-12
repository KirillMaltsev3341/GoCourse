package bst

func (b *Bst) Insert(k int) {
	b.root.insertByNode(k)
}

func (n *node) insertByNode(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			n.right.insertByNode(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.insertByNode(k)
		}
	}
}
