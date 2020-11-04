package main

import "fmt"

func main() {
	// Code for B - Trapezoid Sum
	n := scanInt()
	a := make([][2]float64, n)
	// fmt.Printf("n:%d\n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &a[i][0])
		fmt.Scanf("%f", &a[i][1])
	}
	// fmt.Printf("a:%v\n", a)
	fmt.Println(solve(n, a))
}

func solve(n int, a [][2]float64) int {
	ans := 0.0
	for i := 0; i < n; i++ {
		ans += ((a[i][1] + a[i][0]) / 2) * (a[i][1] - a[i][0] + 1)
	}

	return int(ans)
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
