package main

import "fmt"

func main() {
	// Code for B - Quadruple
	n := scanInt()
	k := scanInt()

	fmt.Println(solve(n, k))
}

func solve(n, k int) int {

	//a + b - c - d = K
	//a: a = K - b + c + d
	//b: b = b = K - a + c + d
	//c: c = -K + a + b - d
	//d: d = -K + a + b - c

	cnt := 0
	minusK := k * (-1)
	for a := 1; a <= n; a++ {
		// fmt.Printf("---a:%d---\n", a)
		for b := 1; b <= n; b++ {
			// fmt.Printf("----b:%d----\n", b)
			for c := 1; c <= n; c++ {
				// fmt.Printf("-----c:%d-----\n", c)

				d := minusK + a + b - c
				// fmt.Printf("***d:%d***\n", d)
				if 1 <= d && d <= n {

					cnt++
				}
			}
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
