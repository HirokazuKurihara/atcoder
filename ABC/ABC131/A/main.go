package main

import (
	"fmt"
)

// A - Security
// https://atcoder.jp/contests/abc131/tasks/abc131_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(ABC131A(s))
}

func ABC131A(s string) string {
	ret := "Good"
	for i := 1; i < 4; i++ {
		if s[i-1] == s[i] {
			ret = "Bad"
		}
	}

	return ret
}
