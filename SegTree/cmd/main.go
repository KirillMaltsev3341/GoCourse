package main

import (
	"fmt"

	"main.go/pkg/segtree"
)

func main() {

	a := []int{5, 4, 2, 3, 5}
	st := segtree.SegTree{}
	st.Build(a)

	fmt.Println(st.Sum(0, 3))
	st.Set(1, 1)
	fmt.Println(st.Sum(0, 3))
	st.Set(3, 1)
	fmt.Println(st.Sum(0, 5))

	fmt.Println("----------")
	st = segtree.SegTree{}
	st.Init(6)
	st.Set(0, 5)
	st.Set(1, 4)
	st.Set(2, 2)
	st.Set(3, 3)
	st.Set(4, 5)
	st.Set(5, 6)

	fmt.Println(st.Sum(0, 6))
	fmt.Println(st.Sum(1, 3))

	st.Set(2, 10)
	fmt.Println(st.Sum(0, 6))
	fmt.Println(st.Sum(1, 3))
}
