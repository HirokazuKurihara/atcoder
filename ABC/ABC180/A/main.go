package main

import (
	"fmt"
)

// A - box
// https://atcoder.jp/contests/abc180/tasks/abc180_a
func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	fmt.Print(ABC180A(n, a, b))
}

func ABC180A(n, a, b int) int {
	return n - a + b
}
