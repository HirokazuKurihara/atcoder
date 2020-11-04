package main

import (
	"fmt"
)

func main() {
	// Code for B - Futon
	h := scanInt()
	w := scanInt()
	slice := make([]string, h)
	for i := 0; i < h; i++ {
		slice[i] = scanString()
	}

	fmt.Println(solve(h, w, slice))
}

func solve(h, w int, slice []string) int {

	ans := 0

	//横の探索
	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			if slice[i][j] == '.' && slice[i][j+1] == '.' {
				ans++
			}
		}
	}

	//縦の探索
	for i := 0; i < h-1; i++ {
		for j := 0; j < w; j++ {
			if slice[i][j] == '.' && slice[i+1][j] == '.' {
				ans++
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
