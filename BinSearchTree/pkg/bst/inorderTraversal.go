package bst

import (
	"fmt"
)

func (b *Bst) InorderTraversal() {
	b.root.inorderTraversalByNode()
	fmt.Printf("\n")
}

func (n *node) inorderTraversalByNode() {
	if n == nil {
		return
	}
	n.left.inorderTraversalByNode()
	fmt.Printf("%d ", n.key)
	n.right.inorderTraversalByNode()
}
