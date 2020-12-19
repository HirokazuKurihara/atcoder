package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for E - Throne
	t := nextInt()
	//fmt.Println(t)

	a := make([][]int, t)
	for i := 0; i < t; i++ {
		a[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			a[i][j] = nextInt() //n, s, k
		}
	}
	//fmt.Println(a)

	ret := solve(t, a)
	for i := 0; i < t; i++ {
		fmt.Println(ret[i])
	}
}

func solve(t int, a [][]int) []int {
	ret := make([]int, t)

	for i := 0; i < t; i++ {
		n := a[i][0]
		s := a[i][1]
		k := a[i][2]

		//TODO:一周でできるかはこれで判断できるが、二周目以降の永遠に座れないの判断をどうするかを検討する必要がある。
		if (n-s)%k == 0 {
			ret[i] = (n - s) / k
		} else {
			ret[i] = -1
		}
	}

	return ret
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
