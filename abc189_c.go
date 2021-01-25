package main

import (
	"bufio"
	"fmt"
	geth "github.com/ethereum/go-ethereum/mobile"
	"os"
	"strconv"
	"strings"
)

// ABC189 C - Mandarin Orange /
// https://atcoder.jp/contests/abc189/tasks/abc189_c
func abc189_c() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	orange_plates := strings.Split(line, " ")

	eatable := 0
	for _, plate := range orange_plates {
		num, _ := strconv.Atoi(plate)

		x <= num
		if 1 <=
	}
}
