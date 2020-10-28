package main

import (
	"fmt"
)

// A - Beginner
// https://atcoder.jp/contests/abc156/tasks/abc156_a
func main() {
	var n, r int
	fmt.Scanf("%d %d", &n, &r)

	fmt.Println(ABC156A(n, r))
}

func ABC156A(n, r int) int {
	if n >= 10 {
		return r
	}

	return r + 100*(10-n)
}
