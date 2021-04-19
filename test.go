package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sk := make([]int, len(s))
	copy(sk, s)
	//sort.Ints(sk)

	fmt.Println(sk)
}
