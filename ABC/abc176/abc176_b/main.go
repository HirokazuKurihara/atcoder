package main

import (
	"fmt"
	"strconv"
)

// B - Multiple of 9
// https://atcoder.jp/contests/abc176/tasks/abc176_b
func main() {
	var n string
	fmt.Scanf("%s", &n)

	fmt.Print(abc176B(n))
}

func abc176B(n string) string {

	sum := 0
	for _, num := range n {
		i, _ := strconv.Atoi(fmt.Sprintf("%c", num))
		sum += i
	}

	if sum%9 == 0 {
		return "Yes"
	} else {
		return "No"
	}
}
