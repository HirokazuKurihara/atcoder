package main

import (
	"fmt"
	"sort"
)

// B - Product Max
// https://atcoder.jp/contests/abc178/tasks/abc178_b
func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	fmt.Print(ABC178B(a, b, c, d))
}

func ABC178B(a, b, c, d int) int {
	slice := make([]int, 0)
	slice = append(slice, a*c)
	slice = append(slice, a*d)
	slice = append(slice, b*c)
	slice = append(slice, b*d)

	//fmt.Printf("%v",slice)
	sort.Ints(slice)
	//fmt.Printf("%v",slice)

	return slice[3]
}
