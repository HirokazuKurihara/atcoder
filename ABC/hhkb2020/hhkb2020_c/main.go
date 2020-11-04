package main

import (
	"fmt"
)

func main() {
	// Code for C - Neq Min
	n := scanInt()
	p := scanInts(n)

	ans := solve(n, p)
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}

func solve(n int, p []int) []int {
	const int32Max int = 2147483647

	ans := make([]int, n)
	avoid := map[int]int{}
	j := 0
	for i := 0; i < n; i++ {
		// fmt.Printf("i:%d\n", i)

		avoid[p[i]] = 1
		// fmt.Printf("avoid:%v\n", avoid)
		for avoid[j] == 1 {
			j++
		}
		ans[i] = j
		// fmt.Printf("ans:%v\n", ans)
	}

	return ans
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
