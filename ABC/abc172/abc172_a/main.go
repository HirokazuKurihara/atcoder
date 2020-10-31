package main

import (
	"fmt"
)

// A - Calc
// https://atcoder.jp/contests/abc172/tasks/abc172_a
func main() {
	var a int
	fmt.Scanf("%d", &a)

	fmt.Print(abc172A(a))
}

func abc172A(a int) int {
	return a + a*a + a*a*a
}
