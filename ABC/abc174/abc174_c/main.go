package main

import (
	"fmt"
)

func main() {
	// Code for C - Repsept
	k := scanInt()
	fmt.Println(solve(k))
}

func solve(k int) int {

	if 7%k == 0 {
		return 1
	}

	num := 7
	for i := 2; i <= k; i++ {
		num = (num*10 + 7) % k
		// fmt.Printf("num:%d, k:%d, num%%k:%d\n", num, k, num%k)
		if num == 0 {
			return i
		}
	}

	return -1
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
