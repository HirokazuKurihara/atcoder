package main

import (
	"fmt"
)

// A - Odds of Oddness
// https://atcoder.jp/contests/abc142/tasks/abc142_a
func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(ABC142A(n))
}

func ABC142A(n int) float64 {
	return float64(n/2+n%2) / float64(n)
}
