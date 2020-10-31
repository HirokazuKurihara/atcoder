package main

import (
	"fmt"
)

// A - Tenki
// https://atcoder.jp/contests/abc139/tasks/abc139_a
func main() {
	var s, t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	fmt.Println(abc139A(s, t))
}

func abc139A(s, t string) int {

	cnt := 0
	for i := 0; i < 3; i++ {
		if s[i] == t[i] {
			cnt++
		}
	}

	return cnt
}
