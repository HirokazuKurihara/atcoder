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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a ...int) int {
	for _, e := range a {
		if e < a[0] {
			a[0] = e
		}
	}
	return a[0]
}

func max(a ...int) int {
	for _, e := range a {
		if e > a[0] {
			a[0] = e
		}
	}
	return a[0]
}

func pow(a, n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= a
	}
	return ret
}
