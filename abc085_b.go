package main

import (
	"fmt"
)

// B - Kagami Mochi
// https://atcoder.jp/contests/abc085/tasks/abc085_b
func abc085_b() {
	var n int
	number_map := map[int]int{}
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scanf("%d", &number)
		number_map[number] = number_map[number] + 1
	}
	fmt.Print(len(number_map))
}
