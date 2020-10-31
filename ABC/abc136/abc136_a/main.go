package main

import (
	"fmt"
)

// A - Transfer
// https://atcoder.jp/contests/abc136/tasks/abc136_a
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(abc136A(a, b, c))
}

func abc136A(a, b, c int) int {
	if c-(a-b) < 0 {
		return 0
	}
	return c - (a - b)
}
