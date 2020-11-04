package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for A - Five Antennas
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()
	e := nextInt()
	k := nextInt()

	fmt.Println(solve(a, b, c, d, e, k))
}

func solve(a, b, c, d, e, k int) string {
	ans := "Yay!"

	minNum := min(a, b, c, d, e)
	maxNum := max(a, b, c, d, e)

	if maxNum-minNum > k {
		ans = ":("
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

func min(a ...int) int {
	for _, e := range a {
		if e < a[0] {
			a[0] = e
		}
	}
	return a[0]
}

func max(a ...int) int {
	for _, e := range a {
		if e > a[0] {
			a[0] = e
		}
	}
	return a[0]
}
