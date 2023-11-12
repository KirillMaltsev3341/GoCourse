package bst

import "fmt"

func (b *Bst) PreorderTraversal() {
	b.root.preorderTraversalByNode()
	fmt.Printf("\n")
}

func (n *node) preorderTraversalByNode() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.key)
	n.left.preorderTraversalByNode()
	n.right.preorderTraversalByNode()
}
