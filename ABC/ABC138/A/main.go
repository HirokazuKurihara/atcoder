package main

import (
	"fmt"
)

// A - Red or Not
// https://atcoder.jp/contests/abc138/tasks/abc138_a
func main() {
	var a int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%s", &s)

	fmt.Println(ABC138A(a, s))
}

func ABC138A(a int, s string) string {
	if a >= 3200 {
		return s
	}

	return "red"
}
