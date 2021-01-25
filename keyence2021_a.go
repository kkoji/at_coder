package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func keyence2021_a() {
	var n int
	fmt.Scanf("%d", &n)

	var str_a, str_b []string
	var a, b []int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")

	str_a = strings.Split(line, " ")
	for i := 0; i < len(str_a); i++ {
		int_a, _ := strconv.Atoi(str_a[i])
		a = append(a, int_a)
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	str_b = strings.Split(line, " ")
	for i := 0; i < len(str_b); i++ {
		int_b, _ := strconv.Atoi(str_b[i])
		b = append(b, int_b)
	}

	max_int := 0
	for i := 0; i < n; i++ {
		for counter := 0; counter <= i; counter++ {
			if max_int < a[counter]*b[counter] {
				max_int = a[counter]*b[counter]
			}
		}
		fmt.Println(max_int)
		max_int = 0
	}
}
