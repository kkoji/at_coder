package main

import "fmt"

// A - Three-Point Shot
// https://atcoder.jp/contests/abc188/tasks/abc188_a
func abc188_a() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	if x > y {
		if x < y + 3 {
			fmt.Print("Yes")
			return
		}
	} else {
		if y < x + 3 {
			fmt.Print("Yes")
			return
		}
	}
	fmt.Print("No")
}
