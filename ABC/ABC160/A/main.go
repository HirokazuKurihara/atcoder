package main

import (
	"fmt"
)

// A - Coffee
// https://atcoder.jp/contests/abc160/tasks/abc160_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(ABC160A(s))
}

// 名前付き戻り値のテスト
func ABC160A(s string) (ret string) {
	ret = "No"
	if s[2] == s[3] && s[4] == s[5] {
		ret = "Yes"
	}

	return
}
