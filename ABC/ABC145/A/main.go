package main

import (
	"fmt"
)

// A - Circle
// https://atcoder.jp/contests/abc145/tasks/abc145_a
func main() {
	var r int
	fmt.Scanf("%d", &r)

	fmt.Println(ABC145A(r))
}

func ABC145A(r int) int {
	return r * r
}
