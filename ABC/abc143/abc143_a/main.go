package main

import (
	"fmt"
)

// A - Curtain
// https://atcoder.jp/contests/abc143/tasks/abc143_a
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(abc143A(a, b))
}

func abc143A(a, b int) int {
	if (a - b*2) <= 0 {
		return 0
	} else {
		return a - b*2
	}
}
