package main

import (
	"fmt"
)

// A - Not
// https://atcoder.jp/contests/abc178/tasks/abc178_a
func main() {
	var x int
	fmt.Scanf("%d", &x)

	fmt.Println(abc178A(x))
}

func abc178A(x int) int {
	if x == 0 {
		return 1
	}
	// 0以外は0を返してしまうが、制限として0or1があるためこれでOK
	return 0
}
