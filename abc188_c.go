package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// C - ABC Tournament
// https://atcoder.jp/contests/abc188/tasks/abc188_c
func abc188_c() {
	var n int
	fmt.Scanf("%d", &n)

	player_rate_map := map[int]int{}
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")

	var player_rates []int
	player_string_rates := strings.Split(line, " ")
	for index, string_rate := range player_string_rates {
		rate, _ := strconv.Atoi(string_rate)
		player_rate_map[rate] = index + 1
		player_rates = append(player_rates, rate)
	}

	var next_player_rates = player_rates
	for true {
		tmp_next_player_rates := []int{}
		for i := 0; i < len(next_player_rates)-1; i += 2 {

			if next_player_rates[i] > next_player_rates[i+1] {
				if len(next_player_rates) == 2 {
					fmt.Print(player_rate_map[next_player_rates[i+1]])
					return
				} else {
					tmp_next_player_rates = append(tmp_next_player_rates, next_player_rates[i])
				}
			} else {
				if len(next_player_rates) == 2 {
					fmt.Print(player_rate_map[next_player_rates[i]])
					return
				} else {
					tmp_next_player_rates = append(tmp_next_player_rates, next_player_rates[i+1])
				}
			}

		}
		next_player_rates = tmp_next_player_rates
	}
}
