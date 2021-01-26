package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ABC081 B - Shift only
// https://atcoder.jp/contests/abc081/tasks/abc081_b
func abc081_b() {
	var n int
	var numbers []int
	fmt.Scanf("%d", &n)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	for _, v := range strings.Split(sc.Text(), " ") {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}

	odd := true
	odd_counter := 0
	for odd {
		var tmp_numbers []int
		for _, number := range numbers {
			if number % 2 != 0 {
				odd = false
				break
			}

			tmp_numbers = append(tmp_numbers, number / 2)
		}
		if odd {
			odd_counter++
			numbers = tmp_numbers
		} else {
			break
		}
	}

	fmt.Print(odd_counter)
}
