package main

import "fmt"

func main() {
	// Code for B - Judge Status Summary
	n := scanInt()
	s := scanStrings(n)
	fmt.Println(solve(n, s))
}

func solve(n int, s []string) string {

	cnt := make([]int, 4)

	for i := 0; i < n; i++ {

		switch s[i] {
		case "AC":
			cnt[0]++
		case "WA":
			cnt[1]++
		case "TLE":
			cnt[2]++
		case "RE":
			cnt[3]++
		}
	}

	ans := fmt.Sprintf("AC x %d\nWA x %d\nTLE x %d\nRE x %d\n", cnt[0], cnt[1], cnt[2], cnt[3])

	return ans
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
