package main

import (
	"fmt"
)

// A - Strings
// https://atcoder.jp/contests/abc149/tasks/abc149_a
func main() {
	var s, t string
	fmt.Scanf("%s %s", &s, &t)

	fmt.Printf("%s", ABC149A(s, t))
}

func ABC149A(k, x string) string {
	return x + k
}
