package main

import "fmt"

// 最大公約数
// ２つ以上の正の整数に共通な約数（公約数）のうち最大のもの
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		// 「 a と b の最大公約数」=「 b と a割るbの余り の最大公約数」となる
		return gcd(b, a%b)
	}
}

// 最小公倍数
// 0でない複数の自然数の公倍数のうち最小の自然数
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

// A - Redundant Redundancy
// https://atcoder.jp/contests/arc110/tasks/arc110_a
func arc110_a() {
	var n int
	fmt.Scanf("%d", &n)
	result := 1
	// 2からnまでの最小公倍数を求める
	for i := 2; i <= n; i++ {
		result = lcm(result, i)
	}
	fmt.Println(result + 1)
}

func arc110_a_2() {
	var n int
	fmt.Scanf("%d", &n)
	k := 1
	for i := 2; i <= n; i++ {
		if k%i != 0 {
			for j := 2; j <= n; j++ {
				if (k*j)%i == 0 {
					k *= j
					break
				}
			}
		}
	}
	fmt.Print(k+1)
}
