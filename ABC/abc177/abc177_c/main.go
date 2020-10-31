package main

import (
	"fmt"
)

// C - Sum of product of pairs
// https://atcoder.jp/contests/abc177/tasks/abc177_c
func main() {
	var n int
	fmt.Scanf("%d", &n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println(abc177C(slice))
}

func abc177C(slice []int) int {
	const MOD = int(1e9) + 7

	cumulativeSumSlice := createCumulativeSumSlice(slice)

	lenSlice := len(slice)
	lenCumulativeSumSlice := len(cumulativeSumSlice)
	sum := 0
	for i := 0; i < lenSlice-1; i++ {
		num := (cumulativeSumSlice[lenCumulativeSumSlice-1] - cumulativeSumSlice[i+1]) % MOD
		sum += slice[i] * num
		sum %= MOD
	}

	return sum
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
