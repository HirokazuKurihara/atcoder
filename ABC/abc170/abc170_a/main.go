package main

import (
	"fmt"
)

// A - Five Variables
// https://atcoder.jp/contests/abc170/tasks/abc170_a
func main() {
	var x [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&x[i])
	}

	fmt.Println(abc170A(x))
}

func abc170A(x [5]int) int {
	ret := 0
	for i := 0; i < 5; i++ {
		if x[i] == 0 {
			ret = i + 1
		}
	}

	return ret
}
