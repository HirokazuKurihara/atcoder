package util

// 累積和のSliceを返す
// param: 数値のSlice（要素数：i)
// return: 累積和のSlice（要素数：i+1）
func createCumulativeSumSlice(slice []int) []int {
	lenSlice := len(slice)
	cumulativeSumSlice := make([]int, lenSlice+1)

	cumulativeSumSlice[0] = 0
	for i := 0; i < lenSlice; i++ {
		cumulativeSumSlice[i+1] = cumulativeSumSlice[i] + slice[i]
	}

	return cumulativeSumSlice
}

// Queue ...
type Queue []int

// pop ...
func (q *Queue) empty() bool {
	return len(*q) == 0
}

// push ...
func (q *Queue) push(i int) {
	*q = append(*q, i)
}

// pop ...
func (q *Queue) pop() (int, bool) {
	if q.empty() {
		return 0, false
	} else {
		res := (*q)[0]
		*q = (*q)[1:]
		return res, true
	}
}

// Stack ...
type Stack []int

// pop ...
func (s *Stack) empty() bool {
	return len(*s) == 0
}

// push ...
func (s *Stack) push(i int) {
	*s = append(*s, i)
}

// pop ...
func (s *Stack) pop() (int, bool) {
	if s.empty() {
		return 0, false
	} else {
		index := len(*s) - 1
		res := (*s)[index]
		*s = (*s)[:index]
		return res, true
	}
}
