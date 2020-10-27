package main

import (
	"fmt"
)

// A - Multiplication 1
// https://atcoder.jp/contests/abc169/tasks/abc169_a
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(ABC169A(a, b))
}

func ABC169A(a, b int) int {
	return a * b
}
