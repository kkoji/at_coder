package main

import (
	"fmt"
	"math"
)

// ABC189 C - Mandarin Orange /
// https://atcoder.jp/contests/abc189/tasks/abc189_c
func abc189_c() {
	// 食べることのできるみかんの数
	ans := 0

	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// 左位置のindex
	for l := 0; l < n; l++ {
		x := a[l]
		// 右位置のindex
		for r := l; r < n; r++ {
			// l, r区間の中の最小値を保存
			x = int(math.Min(float64(x), float64(a[r])))
			// 1つの皿から取れるみかんの数 * 皿数と現在のansで大きい方を保存
			ans = int(math.Max(float64(x*(r-l+1)), float64(ans)))
		}
	}
	fmt.Println(ans)
}
