package main

import (
	"fmt"
)

// A - Rounding
// https://atcoder.jp/contests/abc130/tasks/abc130_a
func main() {
	var x, a int
	fmt.Scanf("%d %d", &x, &a)

	fmt.Println(ABC130A(x, a))
}

func ABC130A(x, a int) int {
	if x >= a {
		return 10
	}
	return 0
}
