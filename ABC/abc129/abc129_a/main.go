package main

import "fmt"

func main() {
	// Code for A - Airplane
	p := scanInt()
	q := scanInt()
	r := scanInt()

	fmt.Println(solve(p, q, r))
}

func solve(p, q, r int) int {
	nums := make([]int, 3)

	nums[0] = p
	nums[1] = q
	nums[2] = r

	ans := 300 //各100が上限
	for i := 0; i < 3-1; i++ {
		for j := i + 1; j < 3; j++ {
			if ans > nums[i]+nums[j] {
				ans = nums[i] + nums[j]
			}
		}
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
