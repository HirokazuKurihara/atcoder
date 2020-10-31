package main

import (
	"fmt"
)

// A - A?C
// https://atcoder.jp/contests/abc166/tasks/abc166_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(abc166A(s))
}

func abc166A(s string) string {
	ret := "abc"
	if s == "abc" {
		ret = "ARC"
	}

	return ret
}
