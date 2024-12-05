package utils

import (
	"strings"
)

func InToGrid(input string) [][]string {
	var l []string = strings.Split(input, "\n")
	l = l[:len(l)-1]
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
