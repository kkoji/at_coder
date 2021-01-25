package main

import (
	"fmt"
)

// ABC081 B - Shift only
// https://atcoder.jp/contests/abc081/tasks/abc081_b
func abc081_b() {
	var N int
	A := make([]int, N)

	for i := range A {
		fmt.Scan(&A[i])
	}

	ans := 0
	for {
		odd := true
		// 次のループで使用する2で割った数を格納する
		for _, v := range A {
			if v%2 != 0 {
				odd = false
			}
		}
		if !odd {
			break
		}

		for i := range A {
			A[i] /= 2
		}
		ans++
	}

	fmt.Print(ans)
}
