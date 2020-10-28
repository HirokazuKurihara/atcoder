package main

import (
	"fmt"
)

// A - We Love Golf
// https://atcoder.jp/contests/abc165/tasks/abc165_a
func main() {
	var k, a, b int
	fmt.Scanf("%d", &k)
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(ABC165A(k, a, b))
}

func ABC165A(k, a, b int) string {
	ret := "NG"
	if a%k == 0 || b%k == 0 {
		ret = "OK"
	}

	if a-(a/k*k) <= k && k <= b-(a/k*k) {
		ret = "OK"
	}

	return ret
}
