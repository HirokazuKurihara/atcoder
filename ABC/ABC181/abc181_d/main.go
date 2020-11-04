package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Code for D - Hachi
	s := scanString()
	fmt.Println(solve(s))
}

// TODO: 未完成。何が悪いのかよくわかってない。
func solve(s string) string {
	lenS := len(s)

	//1桁の場合
	if lenS == 1 {
		if s == "8" {
			return "Yes"
		}

		return "No"
	}

	//2桁の場合
	if lenS == 2 {
		//そのまま判定
		num, _ := strconv.Atoi(s)
		if num%8 == 0 {
			return "Yes"
		}

		//ひっくり返して判定
		num, _ = strconv.Atoi(s[1:2] + s[0:1])
		if num%8 == 0 {
			return "Yes"
		}
		return "No"
	}

	//3桁以上の場合
	if lenS >= 3 {

		//100以上の8の倍数表を作成
		m := map[int]int{}
		for i := 104; i < 1000; i += 8 {
			m[i] = 1
		}

		//下3桁を抽出
		keta1 := s[lenS-1 : lenS]
		keta2 := s[lenS-2 : lenS-1]
		keta3 := s[lenS-3 : lenS-2]

		//下3桁の全パターンを作成
		str321 := keta3 + keta2 + keta1
		str312 := keta3 + keta1 + keta2
		str213 := keta2 + keta1 + keta3
		str231 := keta2 + keta3 + keta1
		str132 := keta1 + keta3 + keta2
		str123 := keta1 + keta2 + keta3
		num321, _ := strconv.Atoi(str321)
		num312, _ := strconv.Atoi(str312)
		num213, _ := strconv.Atoi(str213)
		num231, _ := strconv.Atoi(str231)
		num132, _ := strconv.Atoi(str132)
		num123, _ := strconv.Atoi(str123)
		// fmt.Println(num321)
		// fmt.Println(num312)
		// fmt.Println(num213)
		// fmt.Println(num231)
		// fmt.Println(num132)
		// fmt.Println(num123)

		//下3桁全パターンが倍数表に存在するか確認
		if m[num321] == 1 {
			return "Yes"
		}
		if m[num312] == 1 {
			return "Yes"
		}
		if m[num213] == 1 {
			return "Yes"
		}
		if m[num231] == 1 {
			return "Yes"
		}
		if m[num132] == 1 {
			return "Yes"
		}
		if m[num123] == 1 {
			return "Yes"
		}

		return "No"
	}

	return ""
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
