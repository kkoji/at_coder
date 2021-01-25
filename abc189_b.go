package main

import (
	"fmt"
)

// ABC189 B - Alcoholic
// https://atcoder.jp/contests/abc189/tasks/abc189_b
func abc189_b() {
	var n, x, v, p int
	fmt.Scanf("%d %d", &n, &x)
	x *= 100

	total := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &v, &p)
		total += v * p
		if x < total {
			fmt.Print(i + 1)
			return
		}
	}

	fmt.Print(-1)
}