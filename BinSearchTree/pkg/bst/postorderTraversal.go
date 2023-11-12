package bst

import "fmt"

func (b *Bst) PostorderTraversal() {
	b.root.postorderTraversalByNode()
	fmt.Printf("\n")
}

func (n *node) postorderTraversalByNode() {
	if n == nil {
		return
	}
	n.left.postorderTraversalByNode()
	n.right.postorderTraversalByNode()
	fmt.Printf("%d ", n.key)
}
