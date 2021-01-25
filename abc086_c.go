package main

import (
	"fmt"
	"math"
)

// C - Traveling
// https://atcoder.jp/contests/abc086/tasks/arc089_a
func abc086_c() {
	var n int
	fmt.Scanf("%d", &n)
	// 時刻、x座標, y座標
	t, x, y := []int{0}, []int{0}, []int{0}

	for i := 0; i < n; i++ {
		var i_t, i_x, i_y int
		fmt.Scanf("%d %d %d", &i_t, &i_x, &i_y)
		t = append(t, i_t)
		x = append(x, i_x)
		y = append(y, i_y)
	}

	// 全ての目的地を回れるか
	can := true
	for i := 0; i < n; i++ {
		// 目的地までの所要時間 = 次の目的地の時刻 - 現在の時刻
		dt := t[i+1] - t[i]
		// 現在地から目的地までの距離
		dist := math.Abs(float64(x[i+1]-x[i])) + math.Abs(float64(y[i+1]-y[i]))
		// 所要時間より距離が長い場合は到達不能
		if dt < int(dist) {
			can = false
		}
		// 「その場にとどまることは出来ない」という条件があるため
		// 所要時間と距離が共に奇数、もしくは偶数でない場合は到達不能
		if dt%2 != int(dist)%2 {
			can = false
		}
	}

	if can {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
