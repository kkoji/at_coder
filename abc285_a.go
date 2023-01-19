package main

import "fmt"

// A - Edge Checker 2
// https://atcoder.jp/contests/abc285/tasks/abc285_a
func abc285_a() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if a*2 == b || (a*2)+1 == b {
		fmt.Print("Yes")
		return
	}
	fmt.Print("No")
}
