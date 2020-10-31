package main

import (
	"fmt"
	"math"
)

// A - Password
// https://atcoder.jp/contests/abc140/tasks/abc140_a
func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(abc140A(n))
}

func abc140A(n int) int {
	return int(math.Pow(float64(n), 3))
}
