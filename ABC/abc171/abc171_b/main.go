package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Code for B - Mix Juice
	n := nextInt()
	k := nextInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = nextInt()
	}

	fmt.Println(solve(n, k, nums))
}

func solve(n, k int, nums []int) int {
	ans := 0

	sort.Ints(nums)
	for i := 0; i < k; i++ {
		ans += nums[i]
	}

	return ans
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
