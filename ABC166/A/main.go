package main

import (
	"fmt"
)

// A - A?C
// https://atcoder.jp/contests/abc167/tasks/abc167_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(ABC166A(s))
}

func ABC166A(s string) string {
	ret := "ABC"
	if s == "ABC" {
		ret = "ARC"
	}

	return ret
}
