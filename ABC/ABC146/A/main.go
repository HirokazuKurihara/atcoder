package main

import (
	"fmt"
)

// A - Can't Wait for Holiday
// https://atcoder.jp/contests/abc146/tasks/abc146_a
func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(ABC146A(s))
}

func ABC146A(s string) int {
	weekday := map[string]int{
		"SUN": 7,
		"MON": 6,
		"TUE": 5,
		"WED": 4,
		"THU": 3,
		"FRI": 2,
		"SAT": 1,
	}

	return weekday[s]
}
