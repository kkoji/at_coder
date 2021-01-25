package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// B - Card Game for Two
// https://atcoder.jp/contests/abc088/tasks/abc088_b
func abc088_b() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")

	var a_list []int
	a_string_list := strings.Split(line, " ")
	for _, a_str := range a_string_list {
		a, _ := strconv.Atoi(a_str)
		a_list = append(a_list, a)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a_list)))

	var alice, bob int
	alice_turn := true
	for i := 0; i < len(a_list); i++ {
		if alice_turn {
			alice += a_list[i]
		} else {
			bob += a_list[i]
		}
		alice_turn = !alice_turn
	}

	fmt.Print(alice - bob)
}
