package segtree

func (st *SegTree) Set(i int, v int) {
	st.set(i, v, 0, 0, st.size)
}

func (st *SegTree) set(i int, v int, x int, lx int, rx int) {
	if rx-lx == 1 {
		st.tree[x] = v
		return
	}
	m := (lx + rx) / 2
	if i < m {
		st.set(i, v, 2*x+1, lx, m)
	} else {
		st.set(i, v, 2*x+2, m, rx)
	}
	st.tree[x] = st.tree[2*x+1] + st.tree[2*x+2]
}
