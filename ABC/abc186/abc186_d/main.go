package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Code for D - Sum of difference
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	fmt.Println(solve(n, a))
}

func solve(n int, a []int) int {
	sum := 0

	// absが不要なように降順に変換する
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//fmt.Println(a)
	cumulativeSlice := createCumulativeSumSlice(a)
	//fmt.Println(cumulativeSlice)

	for i := 0; i < n-1; i++ {
		sum += (a[i] * (n - (i + 1))) - (cumulativeSlice[n] - cumulativeSlice[i+1])
		//fmt.Printf("a[i] * (n - (i + 1)):%d\n", a[i]*(n-(i+1)))
		//fmt.Printf("cumulativeSlice[n] - cumulativeSlice[i+1]:%d\n", cumulativeSlice[n]-cumulativeSlice[i+1])
		//fmt.Printf("sum:%d\n", sum)
	}

	return sum
}

/* Template */

var sc *bufio.Scanner

func init() {
	maxBufSize := 1001001
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, maxBufSize), maxBufSize) //Bufferサイズの再定義
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	if !sc.Scan() {
		panic("No more token.")
	}
	num, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic("nextInt(): cannot convert to int: " + sc.Text())
	}
	return num
}

func nextFloat() float64 {
	if !sc.Scan() {
		panic("No more token.")
	}
	numF, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic("nextFloat(): cannot convert to float: " + sc.Text())
	}
	return numF
}

func nextStr() string {
	if !sc.Scan() {
		panic("No more token.")
	}
	return sc.Text()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
