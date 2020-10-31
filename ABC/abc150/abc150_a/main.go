package main

import (
	"fmt"
)

// A - 500 Yen Coins
// https://atcoder.jp/contests/abc150/tasks/abc150_a
func main() {
	var k, x int
	fmt.Scanf("%d %d", &k, &x)

	fmt.Printf("%s", abc150A(k, x))
}

func abc150A(k, x int) string {
	ret := "No"
	if k*500 >= x {
		ret = "Yes"
	}

	return ret
}
