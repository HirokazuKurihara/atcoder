package study

import "fmt"

func fibonacci(n int) int {

	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

var memo = map[int]int{}

func fibonacciMemo(n int) int {

	if memo[n] != 0 {
		return memo[n]
	}

	switch n {
	case 0:
		memo[n] = 0
		return 0
	case 1, 2:
		memo[n] = 1
		return 1
	}
	memo[n] = fibonacciMemo(n-2) + fibonacciMemo(n-1)
	return memo[n]
}

func factorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func hanoi(n int, from, work, dest string) {
	if n >= 2 {
		hanoi(n-1, from, dest, work)
	}

	fmt.Printf("%d を %s　から %s　へ\n", n, from, dest)

	if n >= 2 {
		hanoi(n-1, work, from, dest)
	}
}
