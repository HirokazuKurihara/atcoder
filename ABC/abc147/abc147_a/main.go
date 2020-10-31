package main

import (
	"fmt"
)

// A - Blackjack
// https://atcoder.jp/contests/abc147/tasks/abc147_a
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(abc147A(a, b, c))
}

func abc147A(a, b, c int) string {
	ret := "bust"
	if a+b+c <= 21 {
		ret = "win"
	}

	return ret
}
