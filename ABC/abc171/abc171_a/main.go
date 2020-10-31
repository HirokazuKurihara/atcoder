package main

import (
	"fmt"
)

// A - αlphabet
// https://atcoder.jp/contests/abc171/tasks/abc171_a
func main() {
	var a rune
	fmt.Scanf("%c", &a)

	fmt.Println(abc171A(a))
}

func abc171A(a rune) string {
	if 'A' <= a && a <= 'Z' {
		return "A"
	}

	if 'a' <= a && a <= 'z' {
		return "a"
	}

	// 制限としてa-z,A-Zしか入力はないためこのreturnには到達しえない。
	return ""
}
