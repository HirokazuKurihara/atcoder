package main

import (
	"fmt"
	"strings"
)

// A - Station and Bus
// https://atcoder.jp/contests/abc158/tasks/abc158_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(abc158A(s))
}

func abc158A(s string) string {
	ret := "No"

	if strings.Contains(s, "A") && strings.Contains(s, "B") {
		ret = "Yes"
	}

	return ret
}
