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
