package main

import "fmt"

// A - ABC Preparation
// https://atcoder.jp/contests/abc185/tasks/abc185_a
func abc185_a() {
	var a1, a2, a3, a4 int

	fmt.Scanf("%d %d %d %d", &a1, &a2, &a3, &a4)

	min_int := a1
    if a2 < min_int {
    	min_int = a2
	}
	if a3 < min_int {
		min_int = a3
	}
	if a4 < min_int {
		min_int = a4
	}

	fmt.Print(min_int)
}

