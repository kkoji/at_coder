package main

import "fmt"

// C - 白昼夢
// https://atcoder.jp/contests/abc049/tasks/arc065_a
func abc049_c() {
	strings := []string{"dream", "dreamer", "erase", "eraser"}

	var s string
	fmt.Scanf("%s", &s)
	s = reverse(s)

	// 各文字列の文字を反転させる
	for i := 0; i < len(strings); i++ {
		strings[i] = reverse(strings[i])
	}

	// stringsの文字列でsを作れるかフラグ
	result := true
	// inputされた文字列の文字数分ループ
	for i := 0; i < len(s); {
		// divideできるかフラグ
		match := false

		for j := 0; j < len(strings); j++ {
			str := strings[j]

			if i+len(str) > len(s) {
				continue
			}
			if s[i:i+len(str)] == str {
				match = true
				i += len(str)
				break
			}
		}

		if !match {
			result = false
			break
		}
	}

	if result {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
