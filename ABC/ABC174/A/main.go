package main

import (
	"fmt"
)

// A - Air Conditioner
// https://atcoder.jp/contests/abc174/tasks/abc174_a
func main() {
	var x int
	fmt.Scanf("%d", &x)
	fmt.Print(ABC174A(x))
}

func ABC174A(x int) string {
	ret := "No"
	if x >= 30 {
		ret = "Yes"
	}
	return ret
}
