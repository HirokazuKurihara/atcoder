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
