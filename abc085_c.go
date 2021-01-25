package main

import (
	"fmt"
)

// C - Otoshidama
// https://atcoder.jp/contests/abc085/tasks/abc085_c
func abc085_c() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	for a := 0; a <= n; a++ {
		for b := 0; b <= n; b++ {
			c := n - a - b
			if c < 0 {
				continue
			}
			if a*10000+b*5000+c*1000 == y {
				fmt.Print(a, b, c)
				return
			}
		}
	}

	fmt.Print(-1, -1, -1)
}
