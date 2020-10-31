package main

import (
	"fmt"
)

// A - Next Alphabet
// https://atcoder.jp/contests/abc151/tasks/abc151_a
func main() {
	var c rune
	fmt.Scanf("%c", &c)

	fmt.Printf("%c", abc151A(c))
}

func abc151A(c rune) rune {
	c++
	return c
}
