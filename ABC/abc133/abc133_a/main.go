package main

import (
	"fmt"
)

// A - T or T
// https://atcoder.jp/contests/abc133/tasks/abc133_a
func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	fmt.Println(abc133A(n, a, b))
}

func abc133A(n, a, b int) int {
	if n*a <= b {
		return n * a
	}
	return b
}
