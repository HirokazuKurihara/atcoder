package main

import (
	"fmt"
)

func main() {
	// Code for C - Ubiquity
	n := scanInt()

	fmt.Println(solve(n))
}

func solve(n int) int {
	const MOD = int(1e9) + 7

	k8, k9, k10 := 1, 1, 1
	for i := 0; i < n; i++ {
		k8 *= 8
		k8 %= MOD
		k9 *= 9
		k9 %= MOD
		k10 *= 10
		k10 %= MOD
	}
	ans := (k10 - k9*2 + k8) % MOD
	ans += MOD
	ans %= MOD

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
