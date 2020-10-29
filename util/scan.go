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
		fmt.Scanf("%d", nums[i])
	}
	return nums
}
