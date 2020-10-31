package main

import (
	"fmt"
)

// A - Air Conditioner
// https://atcoder.jp/contests/abc174/tasks/abc174_a
func main() {
	var x int
	fmt.Scanf("%d", &x)
	fmt.Print(abc174A(x))
}

func abc174A(x int) string {
	ret := "No"
	if x >= 30 {
		ret = "Yes"
	}
	return ret
}
