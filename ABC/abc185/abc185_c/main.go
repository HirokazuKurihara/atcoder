package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var mod = int64(1e10 + 6)

func main() {
	// Code for C - Duodecim Ferra
	l := nextInt()
	fmt.Println(solve(l))
}

func solve(l int) *big.Int {

	return combination(l-1, 12-1)
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

func permutation(n int, k int) *big.Int {
	v := big.NewInt(1)
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			k := big.NewInt(int64(n - i))
			v = v.Mul(v, k)
		}
	} else if k > n {
		v = big.NewInt(0)
	}
	return v
}

func factorial(n int) *big.Int {
	return permutation(n, n-1)
}

func combination(n int, k int) *big.Int {
	child := permutation(n, k)
	mother := (factorial(k))
	return child.Div(child, mother)
}
