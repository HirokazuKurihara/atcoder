package main

import (
	"fmt"
)

// A - 9x9
// https://atcoder.jp/contests/abc144/tasks/abc144_a
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(ABC144A(a, b))
}

func ABC144A(a, b int) int {
	if 1 <= a && a <= 9 && 1 <= b && b <= 9 {
		return a * b
	} else {
		return -1
	}
}
