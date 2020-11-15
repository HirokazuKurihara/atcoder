package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for B - Easy Linear Programming
	a := nextInt()
	b := nextInt()
	c := nextInt()
	k := nextInt()
	fmt.Println(solve(a, b, c, k))
}

func solve(a, b, c, k int) int {

	ans := 0
	if a >= k {
		ans += k
		k = 0
	} else {
		ans += a
		k -= a
	}

	if b >= k {
		k = 0
	} else {
		k -= b
	}

	ans -= k

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

func min(a ...int) int {
	for _, e := range a {
		if e < a[0] {
			a[0] = e
		}
	}
	return a[0]
}
