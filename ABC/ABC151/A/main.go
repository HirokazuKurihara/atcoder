package main

import (
	"fmt"
)

// A - Next Alphabet
// https://atcoder.jp/contests/abc151/tasks/abc151_a
func main() {
	var c rune
	fmt.Scanf("%c", &c)

	fmt.Printf("%c", ABC151A(c))
}

func ABC151A(c rune) rune {
	c++
	return c
}
