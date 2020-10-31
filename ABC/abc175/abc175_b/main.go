package main

import (
	"fmt"
	"sort"
)

func main() {
	// Code for B - Making Triangle
	n := scanInt()
	l := scanInts(n)

	fmt.Println(solve(n, l))
}

func solve(n int, l []int) int {

	// fmt.Println(n)
	if n < 3 {
		return 0
	}

	// fmt.Println(l)
	sort.Ints(l)
	// fmt.Println(l)

	cnt := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if l[k] != l[j] && l[i] != l[j] &&
					l[i]+l[j] > l[k] {
					// fmt.Println(l[i], l[j], l[k])
					cnt++
				}
			}
		}
	}

	return cnt
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
