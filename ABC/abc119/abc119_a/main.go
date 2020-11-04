package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	timeLayoutYYYYMMDDHHMMSS = "2006/01/02 15:04:05"
	timeLayoutYYYYMMDD       = "2006/01/02"
	timeLayoutHHMMSS         = "15:04:05"
)

func main() {
	// Code for A - Still TBD
	s := nextStr()
	fmt.Println(solve(s))
}

func solve(s string) string {
	const criteria = "2019/04/30"
	ans := "Heisei"

	tCriteria, _ := time.Parse(timeLayoutYYYYMMDD, criteria)
	t, _ := time.Parse(timeLayoutYYYYMMDD, s)

	if t.Sub(tCriteria) > 0 {
		ans = "TBD"
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
