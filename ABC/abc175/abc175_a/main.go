package main

import (
	"fmt"
)

// A - Rainy Season
// https://atcoder.jp/contests/abc175/tasks/abc175_a
func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Print(abc175A(s))
}

func abc175A(s string) int {
	cnt := 0
	for i := 0; i < 3; i++ {
		if s[i] == 'R' {
			cnt++
		}
	}

	if cnt == 2 {
		if s[1] == 'S' {
			cnt = 1
		}
	}
	return cnt
}
