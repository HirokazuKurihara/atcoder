package main

import (
	"fmt"
	"math"
)

// B - Various distances
// https://atcoder.jp/contests/abc180/tasks/abc180_b
func main() {
	var n int
	fmt.Scanf("%d", &n)
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	manhattan, euclidean, chebyshev := ABC180B(n, slice)
	fmt.Println(manhattan)
	fmt.Printf("%.15f \n", euclidean)
	fmt.Println(chebyshev)

}

func ABC180B(n int, slice []float64) (float64, float64, float64) {
	var manhattan, euclidean, chebyshev float64
	manhattan, euclidean, chebyshev = 0.0, 0.0, 0.0

	//マンハッタン距離
	for i := 0; i < n; i++ {
		manhattan += math.Abs(slice[i])
	}

	//ユークリッド距離
	for i := 0; i < n; i++ {
		euclidean += math.Abs(slice[i]) * math.Abs(slice[i])
	}
	euclidean = math.Sqrt(euclidean)

	//チェビシェフ距離
	for i := 0; i < n; i++ {
		if chebyshev < math.Abs(slice[i]) {
			chebyshev = math.Abs(slice[i])
		}
	}

	return manhattan, euclidean, chebyshev
}
