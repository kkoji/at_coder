package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abc186_b() {
	var h, w, minimum int
	var number_list []int
	fmt.Scanf("%d %d", &h, &w)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < h; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		numbers := strings.Split(text, " ")
		for j, val := range numbers {
			num, _ := strconv.Atoi(val)
			if (i == 0 && j== 0) || num < minimum {
				minimum = num
			}
			number_list = append(number_list, num)
		}
	}

	diff_total := 0
	for _, number := range number_list {
		diff_total += number - minimum
	}
	fmt.Print(diff_total)
}
