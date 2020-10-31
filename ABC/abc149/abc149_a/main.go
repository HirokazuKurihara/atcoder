package main

import (
	"fmt"
)

// A - Strings
// https://atcoder.jp/contests/abc149/tasks/abc149_a
func main() {
	var s, t string
	fmt.Scanf("%s %s", &s, &t)

	fmt.Printf("%s", abc149A(s, t))
}

func abc149A(k, x string) string {
	return x + k
}
