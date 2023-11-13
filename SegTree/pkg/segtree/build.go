package segtree

func (st *SegTree) Build(data []int) {
	st.Init(len(data))
	st.build(data, 0, 0, st.size)
}

func (st *SegTree) build(data []int, x int, lx int, rx int) {
	if rx-lx == 1 {
		if lx < len(data) {
			st.tree[x] = data[lx]
		}
	} else {
		m := (lx + rx) / 2
		st.build(data, 2*x+1, lx, m)
		st.build(data, 2*x+2, m, rx)
		st.tree[x] = st.tree[2*x+1] + st.tree[2*x+2]
	}
}
