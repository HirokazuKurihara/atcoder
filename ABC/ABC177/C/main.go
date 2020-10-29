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
		fmt.Scan(&slice)
	}

	fmt.Println(ABC177C(slice))
}

// TODO:ある程度までできてると思うがまだ完成してない（累積和を用いた解法）
func ABC177C(slice []int) int {
	mod := 1000000007
	//fmt.Printf("slice:%v\n", slice)
	cumulativeSumSlice := createCumulativeSumSlice(slice)
	//fmt.Printf("cumulativeSumSlice:%v\n", cumulativeSumSlice)
	lenSlice := len(slice)
	lenCumulativeSumSlice := len(cumulativeSumSlice)
	//fmt.Printf("lenSlice:%d\n", lenSlice)

	sum := 0
	for i := 0; i < lenSlice-1; i++ {
		//fmt.Printf("i:%d, slice[i]:%d, ", i, slice[i])
		//fmt.Printf("cumulativeSumSlice[lenCumulativeSumSlice-1]:%d, cumulativeSumSlice[i+1]:%d\n", cumulativeSumSlice[lenCumulativeSumSlice-1], cumulativeSumSlice[i+1])
		sum += (slice[i] * (cumulativeSumSlice[lenCumulativeSumSlice-1] - cumulativeSumSlice[i+1])) % mod
	}

	sum %= mod
	return sum
}

func createCumulativeSumSlice(slice []int) []int {
	lenSlice := len(slice)
	cumulativeSumSlice := make([]int, lenSlice+1)

	cumulativeSumSlice[0] = 0
	for i := 0; i < lenSlice; i++ {
		cumulativeSumSlice[i+1] = cumulativeSumSlice[i] + slice[i]
	}

	return cumulativeSumSlice
}
