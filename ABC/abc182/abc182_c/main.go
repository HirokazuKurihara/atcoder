package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Code for C - To 3
	n := nextStr()
	fmt.Println(solve(n))
}

func solve(n string) int {
	lenN := len(n)

	sum := 0
	//3で割れるものを削除（0にする）
	nums := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		nums[i], _ = strconv.Atoi(string(n[i]))
		if nums[i]%3 == 0 {
			nums[i] = 0
		}

		//各桁を足す
		sum += nums[i]
	}

	// fmt.Println("nums:", nums)

	//各桁を足して3で割れるか確認
	if sum%3 == 0 {
		return 0
	}

	//残りの桁同士を足して3で割れるか確認
	for i := 0; i < lenN; i++ {
		if nums[i] == 0 {
			continue
		}
		for j := 0; j < lenN; j++ {
			if nums[j] == 0 {
				continue
			}
			if i == j {
				continue
			}
			if (nums[i]+nums[j])%3 == 0 {
				nums[i] = 0
				nums[j] = 0
				break
			}
		}
	}

	// fmt.Println("nums:", nums)

	cnt := 0
	for i := 0; i < lenN; i++ {
		if nums[i] != 0 {
			cnt++
		}
	}

	if cnt == lenN {
		return -1
	}

	return cnt
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
