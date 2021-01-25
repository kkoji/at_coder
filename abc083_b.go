package main

import "fmt"

// B - Some Sums
// https://atcoder.jp/contests/abc083/tasks/abc083_b
func find_sum_of_digits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func abc083_b() {
	var n, a, b int
	total := 0
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		sum := find_sum_of_digits(i)
		if a <= sum && sum <= b {
			total += i
		}
	}

	fmt.Print(total)
}
