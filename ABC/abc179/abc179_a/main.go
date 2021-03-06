package main

import (
	"fmt"
)

// A - Plural Form
// https://atcoder.jp/contests/abc179/tasks/abc179_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(abc179A(s))
}

func abc179A(s string) string {
	if s[len(s)-1:len(s)] == "s" {
		s = s + "es"
	} else {
		s = s + "s"
	}
	return s
}
