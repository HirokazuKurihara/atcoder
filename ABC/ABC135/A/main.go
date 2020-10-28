package main

import (
	"fmt"
	"strconv"
)

// A - Harmony
// https://atcoder.jp/contests/abc135/tasks/abc135_a
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(ABC135A(a, b))
}

func ABC135A(a, b int) string {
	if (a+b)%2 != 0 {
		return "IMPOSSIBLE"
	}
	return strconv.Itoa((a + b) / 2)
}
