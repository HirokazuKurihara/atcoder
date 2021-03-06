package main

import (
	"fmt"
)

// C - Step
// https://atcoder.jp/contests/abc176/tasks/abc176_c
func main() {
	n := scanInt()
	nums := scanInts(n)

	fmt.Print(abc176C(nums))
}

func abc176C(nums []int) int64 {
	loopMax := len(nums)

	var sum int64
	for i := 0; i < loopMax-1; i++ {
		if nums[i] > nums[i+1] {
			diff := nums[i] - nums[i+1]
			sum += int64(diff)
			nums[i+1] += diff
		}
	}

	return sum
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
