package main

import (
	"fmt"
)

// A - Fifty-Fifty
// https://atcoder.jp/contests/abc132/tasks/abc132_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(abc132A(s))
}

func abc132A(s string) string {
	ret := "No"

	m := map[rune]int{}

	for _, c := range s {
		m[c]++
	}

	if len(m) == 2 {
		ret = "Yes"
	}

	return ret
}
