package main

import (
	"fmt"
)

// A - ABC Swap
// https://atcoder.jp/contests/abc161/tasks/abc161_a
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(ABC161A(a, b, c))
}

func ABC161A(a, b, c int) (int, int, int) {
	var tmp int

	tmp = a
	a = b
	b = tmp

	tmp = a
	a = c
	c = tmp

	return a, b, c
}
