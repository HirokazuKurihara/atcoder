package main

import (
	"fmt"
)

// A - Calc
// https://atcoder.jp/contests/abc172/tasks/abc172_a
func main() {
	var a int
	fmt.Scanf("%d", &a)

	fmt.Print(ABC172A(a))
}

func ABC172A(a int) int {
	return a + a*a + a*a*a
}
