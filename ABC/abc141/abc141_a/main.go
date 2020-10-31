package main

import (
	"fmt"
)

// A - Weather Prediction
// https://atcoder.jp/contests/abc141/tasks/abc141_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(abc141A(s))
}

func abc141A(s string) string {
	weather := map[string]string{
		"Sunny":  "Cloudy",
		"Cloudy": "Rainy",
		"Rainy":  "Sunny",
	}

	return weather[s]
}
