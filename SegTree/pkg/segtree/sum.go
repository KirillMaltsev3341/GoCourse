package segtree

func (st *SegTree) Sum(l int, r int) int {
	return st.sum(l, r, 0, 0, st.size)
}

func (st *SegTree) sum(l int, r int, x int, lx int, rx int) int {
	if l >= rx || lx >= r {
		return 0
	}
	if lx >= l && rx <= r {
		return st.tree[x]
	}
	m := (lx + rx) / 2
	s1 := st.sum(l, r, 2*x+1, lx, m)
	s2 := st.sum(l, r, 2*x+2, m, rx)
	return s1 + s2
}
