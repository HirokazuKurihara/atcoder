package util

import "fmt"

func ScanInt() int {
	var num int
	fmt.Scanf("%d", &num)
	return num
}

func ScanInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	return nums
}

func ScanFloat() float64 {
	var num float64
	fmt.Scanf("%f", &num)
	return num
}

func ScanFloats(n int) []float64 {
	nums := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &nums[i])
	}
	return nums
}

func ScanString() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

func ScanStrings(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &strs[i])
	}
	return strs
}
