package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for A - Brick
	n := nextInt()
	w := nextInt()
	fmt.Println(solve(n, w))
}

func solve(n, w int) int {

	return n / w
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
