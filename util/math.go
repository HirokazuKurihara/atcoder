package util

// フィボナッチ数列の結果を返却する
// param: フィボナッチ数列の求める大きさ
// return: 求められた数式の値
func fib(n int) int {
	fibMemo := make([]int, n)

	fibMemo[0] = 0
	fibMemo[1] = 1
	for i := 2; i <= n; i++ {
		fibMemo[i] = fibMemo[i-1] + fibMemo[i-2]
	}
	return fibMemo[n]
}

// 順列を取得する
// param: 数字配列
// return: 順列の2次元配列
func getPermute(nums []int) [][]int {
	n := fact(len(nums))
	ret := make([][]int, 0, n)
	// fmt.Printf("\n*****\ninit nums:%d\n*****\n", nums)
	permute([]int{}, nums, &ret)
	return ret
}

func permute(pat, nums []int, ret *[][]int) {
	if len(nums) == 0 {
		*ret = append(*ret, pat)
		return
	}
	for i := 0; i < len(nums); i++ {
		// fmt.Printf("-----%d週目-----\n", i)
		// fmt.Printf("ret:%d\n", ret)
		// fmt.Printf("nums[%d]:%d\n", i, nums[i])
		// fmt.Printf("append(pat, nums[i]):%d\n", append(pat, nums[i]))
		// fmt.Printf("nums:%d\n", nums)
		// fmt.Printf("exclude(nums, i):%d\n", exclude(nums, i))
		permute(append(pat, nums[i]), exclude(nums, i), ret)
	}
}

func exclude(nums []int, i int) []int {
	ret := make([]int, 0, len(nums)-1)
	ret = append(ret, nums[:i]...)
	ret = append(ret, nums[i+1:]...)
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func fact(a int) int {
	ans := 1
	for a > 0 {
		ans *= a
		a--
	}
	return ans
}

func min(a ...int) int {
	tmp := a[0]
	for _, e := range a {
		if e < tmp {
			tmp = e
		}
	}
	return tmp
}

func max(a ...int) int {
	tmp := a[0]
	for _, e := range a {
		if e > tmp {
			tmp = e
		}
	}
	return tmp
}

func pow(a, n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= a
	}
	return ret
}
