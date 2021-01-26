package main

import "fmt"

// B - Smartphone Addiction
// https://atcoder.jp/contests/abc185/tasks/abc185_b
func abc185_b() {
	// n バッテリー残量
	// m カフェに訪れた回数
	// t 帰宅時間
	var maxN, n, m, t int

	fmt.Scanf("%d %d %d", &n, &m, &t)
	maxN = n

	var rest_start, rest_end int
	fmt.Scanf("%d %d", &rest_start, &rest_end)
	rest_count := 1
	for i := 1; i <= t; i++ {
		if i > rest_start && i <= rest_end {
			if n < maxN {
				n += 1
			}
		} else {
			n -= 1
		}

		if n == 0 {
			fmt.Println("No")
			return
		}

		// 次回の休憩時間を取得
		if rest_count < m && i > rest_end {
			rest_count += 1
			fmt.Scanf("%d %d", &rest_start, &rest_end)
		}
	}
	fmt.Println("Yes")
}
