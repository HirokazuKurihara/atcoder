package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := nextInt()
	b := nextInt()
	n := nextInt()
	s := nextStr()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = nextInt()
	}

	fmt.Println(solve(a, b, n, s, nums))
}

func solve(a, b, n int, s string, nums []int) int {
	ans := 0

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
