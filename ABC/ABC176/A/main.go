package main

import (
	"fmt"
)

// A - Takoyaki
// https://atcoder.jp/contests/abc176/tasks/abc176_a
func main() {
	var n, x, t int
	fmt.Scanf("%d %d %d", &n, &x, &t)

	fmt.Print(ABC176A(n, x, t))
}

func ABC176A(n, x, t int) int {
	if n%x == 0 {
		return n / x * t
	} else {
		return (n/x + 1) * t
	}
}
