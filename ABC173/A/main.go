package main

import (
	"fmt"
)

// A - Payment
// https://atcoder.jp/contests/abc173/tasks/abc173_a
func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Print(ABC173A(n))
}

func ABC173A(n int) int {
	if n%1000 == 0 {
		return 0
	}
	return 1000 - n%1000
}
