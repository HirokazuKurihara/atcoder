package main

import (
	"fmt"
)

// A - Sheep and Wolves
// https://atcoder.jp/contests/abc164/tasks/abc164_a
func main() {
	var s, w int
	fmt.Scanf("%d %d", &s, &w)

	fmt.Println(abc164A(s, w))
}

func abc164A(s, w int) string {
	ret := "unsafe"
	if s > w {
		ret = "safe"
	}

	return ret
}
