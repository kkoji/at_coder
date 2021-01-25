package main

import "fmt"

// B - Smartphone Addiction
// https://atcoder.jp/contests/abc185/tasks/abc185_b
func abc185_b() {
	// n バッテリー残量
	// m カフェに訪れた回数
	// t 帰宅時間
	var max_n, n, m, t int

	fmt.Scanf("%d %d %d", &n, &m, &t)
	max_n = n

	var rest_start, rest_end int
	fmt.Scanf("%d %d", &rest_start, &rest_end)
	rest_count := 1
	for i := 1; i <= t; i++ {
		if i > rest_start && i <= rest_end {
			if n < max_n {
				n = n + 1
			}
		} else {
			n = n - 1
		}
		fmt.Println("-------------------------")
		fmt.Println("時間: ", i)
		fmt.Println("残量: ", n)
		fmt.Println("休憩: ", rest_start, rest_end)

		if n == 0 {
			fmt.Println("No")
			return
		}

		// 次回の休憩時間を取得
		if rest_count < m && i > rest_end {
			rest_count = rest_count + 1
			fmt.Scanf("%d %d", &rest_start, &rest_end)
		}
	}
	fmt.Println("Yes")
}
