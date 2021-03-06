package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Code for B - Distance
	n := scanInt()
	d := scanInt()

	fmt.Println(solve(n, d))
}

func solve(n, d int) int {

	scanner := bufio.NewReader(os.Stdin)
	x := 0
	y := 0
	cnt := 0
	dd := d * d
	for i := 0; i < n; i++ {
		fmt.Fscan(scanner, &x, &y)
		if x*x+y*y <= dd {
			cnt++
		}
	}

	return cnt
}

func scanInt() int {
	var num int
	fmt.Scanf("%d", &num)
	return num
}

func scanInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	return nums
}

func scanFloat() float64 {
	var num float64
	fmt.Scanf("%f", &num)
	return num
}

func scanFloats(n int) []float64 {
	nums := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &nums[i])
	}
	return nums
}

func scanString() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

func scanStrings(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &strs[i])
	}
	return strs
}
