package main

import (
	"fmt"
)

// A - AC or WA
// https://atcoder.jp/contests/abc152/tasks/abc152_a
func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	fmt.Println(abc152A(n, m))
}

func abc152A(n, m int) string {
	ret := "No"
	if n == m {
		ret = "Yes"
	}

	return ret
}
