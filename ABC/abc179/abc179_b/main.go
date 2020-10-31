package main

import (
	"fmt"
)

// B - Go to Jail
// https://atcoder.jp/contests/abc179/tasks/abc179_b
func main() {
	var n int
	var a, b int
	fmt.Scanf("%d", &n)
	slice := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		slice[i][0], slice[i][1] = a, b
	}

	fmt.Println(abc179B(n, slice))
}

func abc179B(n int, slice [][]int) string {
	cnt := 0
	for i := 0; i < n; i++ {

		if slice[i][0] == slice[i][1] {
			cnt++
			if cnt == 3 {
				return "Yes"
			}
		} else {
			cnt = 0
		}
	}

	return "No"
}
