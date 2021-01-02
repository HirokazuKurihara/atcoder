package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Code for C - 1-SAT
	n := nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextStr()
	}
	fmt.Println(solve(n, s))
}

func solve(n int, s []string) string {

	const ok = "satisfiable"

	sort.Strings(s)
	mStr := map[string]bool{}

	exEnd := 0

	// !の範囲を探索
	for i := 0; i < n; i++ {
		if s[i][0] == '!' {
			mStr[s[i]] = true
			continue
		}
		exEnd = i - 1
		break
	}

	notExStart := exEnd + 1
	notExEnd := n - 1

	for i := notExStart; i <= notExEnd; i++ {

		if mStr["!"+s[i]] {
			return s[i]
		}
	}

	return ok
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
