package main

import (
	"fmt"
)

// A - box
// https://atcoder.jp/contests/abc180/tasks/abc180_a
func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	fmt.Print(abc180A(n, a, b))
}

func abc180A(n, a, b int) int {
	return n - a + b
}
