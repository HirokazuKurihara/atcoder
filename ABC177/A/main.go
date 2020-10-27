package main

import (
	"fmt"
)

// A - Don't be late
// https://atcoder.jp/contests/abc177/tasks/abc177_a
func main() {
	var d, t, s int
	fmt.Scanf("%d %d %d", &d, &t, &s)

	fmt.Print(ABC177A(d, t, s))
}

func ABC177A(d, t, s int) string {
	if d <= t*s {
		return "Yes"
	}
	return "No"
}
