package main

import "fmt"

func main() {
	// Code for C - Typical Stairs
	n := scanInt()
	m := scanInt()
	nums := scanInts(m)

	fmt.Println(solve(n, m, nums))
}

func solve(n, m int, nums []int) int {

	NGList := make([]bool, n+2)
	for i := 0; i < m; i++ {
		NGList[nums[i]] = true
	}

	ans := fib(n, NGList)

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

func fib(n int, NGList []bool) int {
	const MOD int = int(1e9) + 7
	fibMemo := make([]int, n+2)

	//初期値の設定
	fibMemo[0] = 1
	if NGList[1] {
		fibMemo[1] = 0
	} else {
		fibMemo[1] = 1
	}

	for i := 2; i <= n; i++ {
		// fmt.Printf("NGList[%d]:%t\n", i, NGList[i])

		if !NGList[i] {
			//２段連続で壊れた段だった場合は最終段にたどり着けないため終了する
			if NGList[i-1] && NGList[i-2] {
				return 0
			}

			fibMemo[i] = fibMemo[i-1] + fibMemo[i-2]
			fibMemo[i] %= MOD
		} else {
			fibMemo[i] = 0
		}
		// fmt.Printf("fibMemo[%d]:%d\n", i, fibMemo[i])
	}
	return fibMemo[n]
}
