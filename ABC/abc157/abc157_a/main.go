package main

import (
	"fmt"
)

// A - Duplex Printing
// https://atcoder.jp/contests/abc157/tasks/abc157_a
func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(abc157A(n))
}

func abc157A(n int) int {
	ret := 0
	ret += n / 2
	ret += n % 2

	return ret
}
