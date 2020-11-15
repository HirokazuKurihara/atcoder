package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO: 解き途中
func main() {
	// Code for B - Trick or Treat
	n := nextInt()
	k := nextInt()
	dNums := make([]int, n)
	aNums := make([][]int, n)
	fmt.Println(n, k, dNums, aNums)
	for i := 0; i < n; i++ {
		dNums[i] = nextInt()
		fmt.Println(dNums)

		tmpNums := make([]int, dNums[i])
		for j := 0; j < dNums[i]; j++ {
			tmpNums[j] = nextInt()
		}
		aNums[i] = tmpNums
	}

	fmt.Println(n, k, dNums, aNums)

	fmt.Println(solve(n, k, dNums, aNums))
}

func solve(n, k int, dNums []int, aNums [][]int) int {
	m := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < dNums[i]; j++ {
			m[aNums[i][j]]++
		}
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		if m[i] == 0 {
			cnt++
		}
	}

	return cnt
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

func min(a ...int) int {
	for _, e := range a {
		if e < a[0] {
			a[0] = e
		}
	}
	return a[0]
}
