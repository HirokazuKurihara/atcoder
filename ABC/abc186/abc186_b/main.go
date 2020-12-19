package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for B - Blocks on Grid
	h := nextInt()
	w := nextInt()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}

	//fmt.Println(h)
	//fmt.Println(w)
	//fmt.Println(a)
	//for i := 0; i < h; i++ {
	//	for j := 0; j < w; j++ {
	//		fmt.Println(a[i][j])
	//	}
	//}

	fmt.Println(solve(h, w, a))
}

func solve(h, w int, a [][]int) int {
	const int32Max int = 2147483647

	minValue := int32Max
	for i := 0; i < h; i++ {
		tmpMin := min(a[i]...)
		if minValue > tmpMin {
			minValue = tmpMin
		}
	}
	//fmt.Printf("minValue:%d\n", minValue)

	sum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sum += a[i][j] - minValue
			//fmt.Printf("sum += %d - %d   ", a[i][j], minValue)
			//fmt.Printf("i:%d,j:%d,sum:%d\n", i, j, sum)
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
	tmp := a[0]
	for _, e := range a {
		if e < tmp {
			tmp = e
		}
	}
	return tmp
}
