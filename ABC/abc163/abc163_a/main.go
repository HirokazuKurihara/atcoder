package main

import (
	"fmt"
	"math"
)

// A - Circle Pond
// https://atcoder.jp/contests/abc163/tasks/abc163_a
func main() {
	var r int
	fmt.Scanf("%d", &r)

	fmt.Println(abc163A(r))
}

func abc163A(r int) float64 {
	return float64(2*r) * math.Pi
}
