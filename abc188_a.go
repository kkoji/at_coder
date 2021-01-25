package main

import "fmt"

func abc188_a() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	if x > y {
		if x < y + 3 {
			fmt.Print("Yes")
			return
		}
	} else {
		if y < x + 3 {
			fmt.Print("Yes")
			return
		}
	}
	fmt.Print("No")
}
