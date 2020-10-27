package main

import (
	"fmt"
)

// C - A x B + C
// https://atcoder.jp/contests/abc179/tasks/abc179_c
func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(ABC179C(n))
}

func ABC179C(n int) int {
	cnt := 0
	for a := 1; a < n; a++ {
		for b := 1; b < n; b++ {
			if a*b < n {
				cnt++
			} else {
				break
			}
		}
	}

	return cnt
}
