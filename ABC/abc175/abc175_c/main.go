package main

import (
	"fmt"
)

func main() {
	// Code for C - Walking Takahashi
	x := scanInt()
	k := scanInt()
	d := scanInt()

	fmt.Print(solve(x, k, d))
}

func solve(x, k, d int) int {

	x = abs(x)
	cnt := min(k, x/d)

	x -= d * cnt
	k -= cnt

	if k%2 == 1 {
		x = d - x
	}

	return x
}

func scanInt() int {
	var num int
	fmt.Scanf("%d", &num)
	return num
}

func scanInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	return nums
}

func scanFloat() float64 {
	var num float64
	fmt.Scanf("%f", &num)
	return num
}

func scanFloats(n int) []float64 {
	nums := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &nums[i])
	}
	return nums
}

func scanString() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

func scanStrings(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &strs[i])
	}
	return strs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a ...int) int {
	for _, e := range a {
		if e < a[0] {
			a[0] = e
		}
	}
	return a[0]
}
