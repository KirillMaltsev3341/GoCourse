package segtree

func (st *SegTree) Init(n int) {
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.tree = make([]int, 2*st.size-1)
}
