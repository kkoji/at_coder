package main

import "fmt"

// B - Longest Uncommon Prefix
// https://atcoder.jp/contests/abc285/tasks/abc285_b
func abc285_b() {
	var n int
	var s string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)

	for i := 1; i < n; i++ {
		for j := 1; j <= n; j++ {
			// n（文字数）の範囲を超えた場合
			if i+j > n {
				fmt.Println(j - 1)
				break
			}
			// j番目の文字とj+i番目の文字が同一である場合
			if s[j-1] == s[j+i-1] {
				fmt.Println(j - 1)
				break
			}
		}
	}
}
