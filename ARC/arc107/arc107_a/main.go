package main

import "fmt"

func main() {
	// Code for A - Simple Math
	a := scanInt()
	b := scanInt()
	c := scanInt()

	fmt.Println(solve(a, b, c))
}

func solve(A, B, C int) int {
	const MOD = 998244353

	cSum := 0
	if C%2 == 0 {
		cSum = ((1 + C) * (C / 2)) % MOD
	} else {
		cSum = ((1+C)*(C/2) + (C/2 + 1)) % MOD
	}
	// fmt.Println(cSum)

	bSum := 0
	if B%2 == 0 {
		bSum = ((1 + B) * (B / 2)) % MOD
	} else {
		bSum = ((1+B)*(B/2) + (B/2 + 1)) % MOD
	}
	// fmt.Println(bSum)

	aSum := 0
	if A%2 == 0 {
		aSum = ((1 + A) * (A / 2)) % MOD
	} else {
		aSum = ((1+A)*(A/2) + (A/2 + 1)) % MOD
	}
	// fmt.Println(aSum)

	ans := ((aSum * bSum) % MOD * cSum) % MOD
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
