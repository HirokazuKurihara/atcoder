package main

import (
	"fmt"
)

// A - Round One
// https://atcoder.jp/contests/abc148/tasks/abc148_a
func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	fmt.Println(abc148A(a, b))
}

func abc148A(a, b int) int {
	return 6 - a - b
}
