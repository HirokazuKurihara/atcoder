package main

import (
	"fmt"
	"strings"
)

// A - Registration
// https://atcoder.jp/contests/abc167/tasks/abc167_a
func main() {
	var s, t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	fmt.Println(ABC167A(s, t))
}

func ABC167A(s, t string) string {
	ret := "No"
	if strings.Index(t, s) == 0 {
		if len(s)+1 == len(t) {
			ret = "Yes"
		}
	}

	return ret
}
