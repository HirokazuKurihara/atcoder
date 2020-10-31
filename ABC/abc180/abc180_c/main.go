package main

import (
	"fmt"
	"math"
	"sort"
)

// C - Cream puff
// https://atcoder.jp/contests/abc180/tasks/abc180_c
func main() {
	var n int
	fmt.Scanf("%d", &n)

	intSlice := abc180C(n)
	for _, i := range intSlice {
		fmt.Println(i)
	}
}

func abc180C(n int) []int {
	var m = map[int]int{}
	max := int(math.Sqrt(float64(n)))

	for i := 1; i <= max; i++ {
		if n%i == 0 {
			m[i] = i
			m[n/i] = n / i
		}
	}

	var intSlice []int
	for i := range m {
		intSlice = append(intSlice, i)
	}
	sort.Ints(intSlice)

	return intSlice
}
