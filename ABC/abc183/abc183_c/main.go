package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for C - Travel
	n := nextInt()
	k := nextInt()
	numss := make([][]int, n)
	for i := 0; i < n; i++ {
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			nums[j] = nextInt()
		}
		numss[i] = nums
	}
	// fmt.Println(numss)
	fmt.Println(solve(n, k, numss))
}

func solve(n, k int, numss [][]int) int {
	ans := 0
	size = n
	fmt.Println(size)

	for i := 1; i < n; i++ {
		if kumiawase(0, i) == k {
			ans++
		}
	}

	return ans
}

var size int
var memo = map[string]int{}

func kumiawase(i, j int) int {
	ijStr := fmt.Sprintf("%d%d", i, j)

	fmt.Println(i, j)
	fmt.Println(memo[ijStr])
	if memo[ijStr] != 0 {

		return memo[ijStr]
	}

	// ans := kumiawase((i+1)%size, (j+1)%size) + kumiawase(i%size, j%size)
	ans := 1
	memo[ijStr] = ans
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

func fact(a int) int {
	ans := 1
	for a > 0 {
		ans *= a
		a--
	}
	return ans
}
