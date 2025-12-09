package main

import (
	"strconv"
	"strings"
)

func num_lines(s string) []int {
	res1 := make([]int, 0)
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	for _, v := range lines {
		in, _ := strconv.Atoi(v)
		res1 = append(res1, in)
	}
	return res1

}
