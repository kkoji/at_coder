package main

import (
	"fmt"
)

// C - Bowls and Dishes
// https://atcoder.jp/contests/abc190/tasks/abc190_c

func abc190_c() {
	// N = 皿の数
	// M = 条件の数
	var N, M int
	fmt.Scan(&N, &M)

	// A, B = 満たすべき2つの条件となる皿ペア
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&A[i], &B[i])
	}

	var K int
	fmt.Scan(&K)

	// 各人が選択可能な皿C[i] or D[i]
	C := make([]int, K)
	D := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&C[i], &D[i])
	}

	ans := 0
	// K人が2つの皿のいずれかを選ぶので、2のk乗通りのパターンを試す
	for bit := 0; bit < 1<<K; bit++ {
		dish := make([]int, N+1)
		// 組み合わせごとにK人数選択する
		for i := 0; i < K; i++ {
			// i人目の人がどっちを選択肢たか
			// ビットを一つずつ右にずらして、一番右のビットが1であるかを判定
			// 要するに全てのビットをチェックしている
			if bit >> i & 1 == 1 {
				// Cを選択肢たケース
				dish[C[i]]++
			} else {
				// Dを選択肢たケース
				dish[D[i]]++
			}
		}
		now := 0
		// 両方の皿に1以上のボールがあるものだけカウント
		for i := 0; i < M; i++ {
			if dish[A[i]] == 0 {
				continue
			}
			if dish[B[i]] == 0 {
				continue
			}
			now++
		}
		// 条件を満たす最大値を取得
		if now > ans {
			ans = now
		}
	}
	fmt.Println(ans)
}
