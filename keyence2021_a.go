package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A - Two Sequences 2
// https://atcoder.jp/contests/keyence2021/tasks/keyence2021_a
func keyence2021A() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	for i := range b {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
	}
	// 答えの数列を格納する変数
	c := make([]int, n)
	c[0] = a[0] * b[0]
	fmt.Println(c[0])

	maxA := a[0]
	for i := 1; i < n; i++ {
		if maxA < a[i] {
			// a[0]からa[i]までの数列での最大値を格納
			maxA = a[i]
		}
		// 1つ前のcと、aの最大値とbの現在位置の乗算で大きい方を格納
		// ※常に1つ前のc以上の数値になる
		if c[i-1] < maxA*b[i] {
			c[i] = maxA * b[i]
		} else {
			c[i] = c[i-1]
		}

		fmt.Println(c[i])
	}
}
