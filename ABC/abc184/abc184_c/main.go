package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for C - Super Ryuma
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()
	fmt.Println(solve(a, b, c, d))
}

func solve(a, b, c, d int) int {

	//マイナスを０地点に移動させる
	if a < 0 {
		c -= a
		a = 0
	}
	if b < 0 {
		d -= b
		b = 0
	}
	if c < 0 {
		a -= c
		c = 0
	}
	if d < 0 {
		b -= d
		d = 0
	}

	//同じ位置
	if a == c && b == d {
		return 0
	}

	//1歩で移動できる範囲内
	if a+b == c+d || a-b == c-d || abs(a-c)+abs(b-d) <= 3 {
		return 1
	}

	//X軸同じまで移動後に1歩で移動できる範囲
	sabun := c - a
	if d-3 <= b+sabun && b+sabun <= d+3 || d-3 <= b-sabun && b-sabun <= d+3 {
		return 2
	}

	//斜めに移動すれば同じ斜線上に入る位置
	if (abs(a-c)+abs(b-d))%2 == 0 {
		return 2
	}

	//それ以外は3歩になる
	return 3
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
