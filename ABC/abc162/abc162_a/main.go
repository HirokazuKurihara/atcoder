package main

import (
	"fmt"
	"strings"
)

// A - Lucky 7
// https://atcoder.jp/contests/abc162/tasks/abc162_a
func main() {
	var n string
	fmt.Scanf("%s", &n)

	fmt.Println(abc162A(n))
}

func abc162A(n string) string {
	ret := "No"
	if strings.Contains(n, "7") {
		ret = "Yes"
	}
	return ret
}
