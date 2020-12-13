package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for B - Smartphone Addiction
	n := nextInt()
	m := nextInt()
	t := nextInt()
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
	}
	fmt.Println(solve(n, m, t, a, b))
}

func solve(n, m, t int, a, b []int) string {

	battery := n

	// 1回目
	battery -= a[0]
	//fmt.Println(battery)
	if battery <= 0 {
		return "No"
	}
	battery += b[0] - a[0]
	if battery >= n {
		battery = n
	}
	//fmt.Println(battery)

	//　2〜m-1回目
	for i := 1; i < m; i++ {
		battery -= a[i] - b[i-1]
		if battery <= 0 {
			return "No"
		}
		//fmt.Println(battery)
		battery += b[i] - a[i]
		if battery >= n {
			battery = n
		}
		//fmt.Println(battery)
	}

	// m回目
	battery -= t - b[m-1]
	//fmt.Println(battery)
	if battery <= 0 {
		return "No"
	}

	return "Yes"
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
