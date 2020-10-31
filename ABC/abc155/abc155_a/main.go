package main

import (
	"fmt"
)

// A - Poor
// https://atcoder.jp/contests/abc155/tasks/abc155_a
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(abc155A(a, b, c))
}

func abc155A(a, b, c int) string {
	ret := "Yes" //かわいそう

	if a == b && b == c {
		ret = "No"
	}

	if a != b && b != c && a != c {
		ret = "No"
	}

	return ret
}
