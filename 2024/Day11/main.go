package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser() []int {
	nums := strings.Split(strings.TrimSuffix(inputDay, "\n"), " ")
	res := make([]int, 0)
	for i := range nums {
		n1, _ := strconv.Atoi(nums[i])
		res = append(res, n1)
	}
	return res
}

func cut_number(s string) (int, int) {
	tab := strings.Split(s, "")
	n := len(tab)
	n1_s := ""
	n2_s := ""
	for i := range n / 2 {
		n1_s += tab[i]
		n2_s += tab[i+(n/2)]
	}
	n1, _ := strconv.Atoi(n1_s)
	n2, _ := strconv.Atoi(n2_s)
	return n1, n2
}

func blink(block int, turn int, seen map[[2]int]int) int {
	if seen[[2]int{block, turn}] != 0 {
		return seen[[2]int{block, turn}]
	}
	var res int = 0
	n_str := strings.Split(strconv.Itoa(block), "")
	if turn == 0 {
		return 1
	}
	if block == 0 {
		res += blink(1, turn-1, seen)
	} else if len(n_str)%2 == 0 {
		n1, n2 := cut_number(strconv.Itoa(block))
		res += blink(n1, turn-1, seen)
		res += blink(n2, turn-1, seen)
	} else {
		res += blink(block*2024, turn-1, seen)
	}
	seen[[2]int{block, turn}] = res
	return res
}

func Part1(turn int, blocks []int) int {
	res := 0
	seen := make(map[[2]int]int)
	for _, v := range blocks {
		res += blink(v, turn, seen)
	}
	return res
}
func Part2() {
	return //HAHAHAHA
}

func main() {
	fmt.Println(parser())
	fmt.Println("PART1:", Part1(25, parser()))
	fmt.Println("PART 2:", Part1(75, parser()))
}
