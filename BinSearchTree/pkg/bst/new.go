package bst

func New(k int) Bst {
	return Bst{root: &node{key: k}}
}
