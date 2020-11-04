package main

import (
	"fmt"
)

func main() {
	// Code for C - Collinearity
	n := scanInt()
	a := make([][2]float64, n)
	// fmt.Printf("n:%d\n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &a[i][0])
		fmt.Scanf("%f", &a[i][1])
	}
	// fmt.Printf("a:%v\n", a)

	fmt.Println(solve(n, a))
}

func solve(n int, a [][2]float64) string {
	// ∣AB∣+∣AC∣=∣BC∣
	// or
	// ∣AB∣+∣BC∣=∣AC∣
	// or
	// ∣AC∣+∣BC∣=∣AB∣

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := a[i][0], a[i][1]
				x2, y2 := a[j][0], a[j][1]
				x3, y3 := a[k][0], a[k][1]

				x1 -= x3
				x2 -= x3
				y1 -= y3
				y2 -= y3
				if x1*y2 == x2*y1 {
					// fmt.Println(a[i][0], a[i][1])
					// fmt.Println(a[j][0], a[j][1])
					// fmt.Println(a[k][0], a[k][1])
					return "Yes"
				}
			}
		}
	}

	return "No"
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
