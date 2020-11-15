package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Code for B - Almost GCD
	n := nextInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = nextInt()
	}
	fmt.Println(solve(n, nums))
}

func solve(n int, nums []int) int {

	sort.Ints(nums)
	// fmt.Println(nums)
	ans := 0
	maxCnt := 0
	for i := 2; i <= nums[n-1]; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if nums[j]%i == 0 {
				cnt++
			}
		}

		if maxCnt <= cnt {
			maxCnt = cnt
			ans = i

			// fmt.Println("maxCnt:", maxCnt, "ans:", ans)
		}
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
