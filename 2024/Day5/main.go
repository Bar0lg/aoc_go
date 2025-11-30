package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var InputDay string

func to_int(t [][]string) [][]int {
	var res [][]int = make([][]int, 0)
	for i, v := range t {
		res = append(res, make([]int, 0))
		for _, s := range v {
			var n1, _ = strconv.Atoi(s)
			res[i] = append(res[i], n1)
		}
	}
	return res
}

func parser() ([][]int, [][]int) {
	var cutted []string = strings.Split(InputDay, "\n\n")
	var orderL []string = strings.Split(cutted[0], "\n")
	var presentPage []string = strings.Split(cutted[1], "\n")
	presentPage = presentPage[:len(presentPage)-1]
	var order_res [][]string = make([][]string, 0)
	var pages_res [][]string = make([][]string, 0)
	for _, v := range orderL {
		order_res = append(order_res, strings.Split(v, "|"))
	}
	for _, v := range presentPage {
		pages_res = append(pages_res, strings.Split(v, ","))
	}
	return to_int(order_res), to_int(pages_res)
}

func in_tab(elem int, t []int) bool {
	var res bool = false
	for _, v := range t {
		res = res || (v == elem)
	}
	return res
}

func part1() int {
	var res int
	pageOr, pagePre := parser()
	for _, manual := range pagePre {
		var seen []int = make([]int, 0)
		var valid bool = true
		for _, page := range manual {
			for _, rule := range pageOr {
				if rule[1] == page {
					if in_tab(rule[0], manual) && !in_tab(rule[0], seen) {
						valid = false
					}
				}
			}
			seen = append(seen, page)
		}
		if valid {
			res += manual[len(manual)/2]
		}
	}
	return res
}

func repair(page_O [][]int, broken []int) []int {
	var res []int = make([]int, 0)
	var index int = 0
	for len(res) != len(broken) {
		var valid bool = true
		if !in_tab(broken[index], res) {
			for _, rule := range page_O {
				if rule[1] == broken[index] {
					if in_tab(rule[0], broken) {
						if !in_tab(rule[0], res) {
							valid = false
						}
					}
				}
			}
			if valid {
				res = append(res, broken[index])
			}
		}
		index = (index + 1) % len(broken)
	}
	return res
}

func part2() int {
	var res int
	pageOr, pagePre := parser()
	for _, manual := range pagePre {
		var seen []int = make([]int, 0)
		var invalid = false
		for _, page := range manual {
			for _, rule := range pageOr {
				if rule[1] == page {
					if in_tab(rule[0], manual) && !in_tab(rule[0], seen) {
						r := repair(pageOr, manual)
						res += r[len(r)/2]
						invalid = true
						break
					}
				}
			}
			if invalid {
				break
			}
			seen = append(seen, page)
		}
	}
	return res
}
func main() {
	//x,y := parser()
	//fmt.Println(repair(x,y[4]))
	fmt.Println(part2())
}
