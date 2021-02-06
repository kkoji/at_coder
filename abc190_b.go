package main

import "fmt"

// B - Magic 3
// https://atcoder.jp/contests/abc190/tasks/abc190_b
func abc190_b() {
	var N, S, D int
	fmt.Scan(&N, &S, &D)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		if X < S && Y > D {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
