package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for B - Gentle Pairs
	n := nextInt()
	xy := make([][]float64, n)
	for i := 0; i < n; i++ {
		tmp := make([]float64, 2)
		tmp[0] = nextFloat()
		tmp[1] = nextFloat()
		xy[i] = tmp
	}
	//fmt.Println(xy)

	fmt.Println(solve(n, xy))
}

func solve(n int, xy [][]float64) int {

	sum := 0

	//xy[i][0] = x1
	//xy[i][1] = y1
	//xy[j][0] = x2
	//xy[j][1] = y2

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			//分母が0にならないかチェック
			if (xy[j][0] - xy[i][0]) == 0 {
				continue
			}

			//傾きをチェック(-1以上、1以下）
			katamuki := (xy[j][1] - xy[i][1]) / (xy[j][0] - xy[i][0])
			if -1 <= katamuki && katamuki <= 1 {
				sum++
			}
		}
	}

	return sum
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
