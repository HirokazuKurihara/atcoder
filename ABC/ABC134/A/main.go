package main

import (
	"fmt"
)

// A - Dodecagon
// https://atcoder.jp/contests/abc134/tasks/abc134_a
func main() {
	var r int
	fmt.Scanf("%d", &r)

	fmt.Println(ABC134A(r))
}

func ABC134A(r int) int {
	return 3 * r * r
}
