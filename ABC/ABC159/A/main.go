package main

import (
	"fmt"
)

// A - The Number of Even Pairs
// https://atcoder.jp/contests/abc159/tasks/abc159_a
func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	fmt.Println(ABC159A(n, m))
}

func ABC159A(n, m int) int {
	ret := 0
	ret += combination(n, 2)
	ret += combination(m, 2)

	return ret
}

func combination(n, r int) int {
	if n < r {
		return 0
	}

	bo, si := 1, 1
	for i := 0; i < r; i++ {
		bo *= n - i
	}
	for i := 0; i < r; i++ {
		si *= i + 1
	}

	return bo / si
}
