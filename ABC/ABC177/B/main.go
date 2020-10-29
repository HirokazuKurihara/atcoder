package main

import (
	"fmt"
)

// B - Substring
// https://atcoder.jp/contests/abc177/tasks/abc177_b
func main() {
	var s, t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	fmt.Println(ABC177B(s, t))
}

func ABC177B(s, t string) int {
	lenS := len(s)
	lenT := len(t)
	//fmt.Printf("lenS:%d, lenT:%d\n", lenS, lenT)

	ret := lenT
	for i := 0; i <= lenS-lenT; i++ {
		cnt := 0
		//fmt.Printf("---%d週目---\n", i+1)
		for j := 0; j < lenT; j++ {
			//fmt.Printf("s[%d]:%c, t[%d]:%c\n", i+j, s[i+j], j, t[j])
			if s[i+j] != t[j] {
				cnt++
			}
		}
		//fmt.Printf("%d週目結果:%d\n", i+1, cnt)
		if ret > cnt {
			ret = cnt
		}
	}
	//fmt.Printf("---ret:%d---\n", ret)
	return ret
}
