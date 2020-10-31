package main

import (
	"fmt"
	"math"
)

func main() {
	// Code for B - Balance
	n := scanInt()
	nums := scanInts(n)

	fmt.Println(solve(nums))
}

func solve(nums []int) int {
	loopMax := len(nums)
	for i := 0; i < loopMax; i++ {
		nums[i] = int(math.Abs(float64(nums[i])))
	}

	// fmt.Printf("nums:%v\n", nums)

	CuSumSlice := createCumulativeSumSlice(nums)

	// fmt.Printf("CuSumSlice:%v\n", CuSumSlice)

	ans := 2147483647
	for i := 1; i < loopMax; i++ {
		// Tの左側 - Tの右側
		// T = i
		ret := int(math.Abs(float64((CuSumSlice[i] - CuSumSlice[0]) - (CuSumSlice[loopMax] - CuSumSlice[i]))))
		fmt.Println(ret)
		if ans > ret {
			ans = ret
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

// 累積和のSliceを返す
// param: 数値のSlice（要素数：i)
// return: 累積和のSlice（要素数：i+1）
func createCumulativeSumSlice(slice []int) []int {
	lenSlice := len(slice)
	cumulativeSumSlice := make([]int, lenSlice+1)

	cumulativeSumSlice[0] = 0
	for i := 0; i < lenSlice; i++ {
		cumulativeSumSlice[i+1] = cumulativeSumSlice[i] + slice[i]
	}

	return cumulativeSumSlice
}
