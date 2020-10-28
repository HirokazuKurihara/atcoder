package main

import (
	"fmt"
	"sort"
)

// A - +-x
// https://atcoder.jp/contests/abc137/tasks/abc137_a
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(ABC137A(a, b))
}

func ABC137A(a, b int) int {
	ints := make([]int, 3)

	ints[0] = a + b
	ints[1] = a - b
	ints[2] = a * b

	sort.Ints(ints)

	return ints[2]
}
