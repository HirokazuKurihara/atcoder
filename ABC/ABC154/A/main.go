package main

import (
	"fmt"
)

// A - Remaining Balls
// https://atcoder.jp/contests/abc154/tasks/abc154_a
func main() {
	var s, t, u string
	var a, b int
	fmt.Scanf("%s %s", &s, &t)
	fmt.Scanf("%d %d", &a, &b)
	fmt.Scanf("%s", &u)

	fmt.Println(ABC154A(s, t, a, b, u))
}

func ABC154A(s, t string, a, b int, u string) (int, int) {
	if s == u {
		a--
	}
	if t == u {
		b--
	}

	return a, b

}
