package main

import (
	"fmt"
)

// A - Serval vs Monster
// https://atcoder.jp/contests/abc153/tasks/abc153_a
func main() {
	var h, a int
	fmt.Scanf("%d %d", &h, &a)

	fmt.Println(abc153A(h, a))
}

func abc153A(h, a int) int {
	ret := 0
	if h%a == 0 {
		ret = h / a
	} else {
		ret = h/a + 1
	}

	return ret
}
