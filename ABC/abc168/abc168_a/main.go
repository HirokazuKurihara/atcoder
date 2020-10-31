package main

import (
	"fmt"
)

// A - ∴ (Therefore)
// https://atcoder.jp/contests/abc168/tasks/abc168_a
func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(abc168A(n))
}

func abc168A(n int) string {
	n = n % 10

	switch n {
	//hon
	case 2, 4, 5, 7, 9:
		return "hon"
	//pon
	case 0, 1, 6, 8:
		return "pon"
	//bon
	case 3:
		return "bon"
	default:
		return ""
	}
}
