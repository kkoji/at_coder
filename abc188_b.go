package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abc188_b() {
	var n int
	var a, b []int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	a_string := strings.Split(line, " ")
	for _, a_string_item := range a_string {
		a_int_item, _ := strconv.Atoi(a_string_item)
		a = append(a, a_int_item)
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	b_string := strings.Split(line, " ")
	for _, b_string_item := range b_string {
		b_int_item, _ := strconv.Atoi(b_string_item)
		b = append(b, b_int_item)
	}

	inner_product := 0
	for i := 0; i < n; i++ {
		inner_product += a[i] * b[i]
	}
	if inner_product == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
