package main

import "fmt"

// A - Very Very Primitive Game
// https://atcoder.jp/contests/abc190/tasks/abc190_a
func abc190_a() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a < b {
		fmt.Println("Aoki")
	} else if a > b  {
		fmt.Println("Takahashi")
	} else {
		if c == 0 {
			fmt.Println("Aoki")
		} else {
			fmt.Println("Takahashi")
		}
	}
}
