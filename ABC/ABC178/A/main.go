package main

import (
	"fmt"
)

// A - Not
// https://atcoder.jp/contests/abc178/tasks/abc178_a
func main() {
	var x int
	fmt.Scanf("%d", &x)

	fmt.Println(ABC178A(x))
}

func ABC178A(x int) int {
	if x == 0 {
		return 1
	}
	// 0以外は0を返してしまうが、制限として0or1があるためこれでOK
	return 0
}
