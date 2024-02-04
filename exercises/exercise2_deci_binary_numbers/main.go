package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func minPartitions(n string) int {
	var max = 0
	for _, ch := range strings.Split(n, "") {
		number, err := strconv.Atoi(ch)
		if err != nil {
			panic(err)
		}
		if number > max {
			max = number
			if number == 9 {
				break
			}
		}
	}
	return max
}

func minPartitions2(n string) int {
	n_arr := strings.Split(n, "")
	for i := 9; i >= 1; i-- {
		s_n := strconv.Itoa(i)
		if slices.Contains(n_arr, s_n) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(minPartitions2("32"))                   // 3
	fmt.Println(minPartitions2("82734"))                // 8
	fmt.Println(minPartitions2("27346209830709182346")) // 9
}
