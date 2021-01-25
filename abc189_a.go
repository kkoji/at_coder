package main

import (
	"fmt"
)

// ABC189 A - Slot
// https://atcoder.jp/contests/abc189/tasks/abc189_a
func abc189_a() {
	var c string
	fmt.Scanf("%s", &c)

	if c[0:1] == c[1:2] && c[1:2] == c[2:3] {
		fmt.Print("Won")
	} else {
		fmt.Print("Lost")
	}
}