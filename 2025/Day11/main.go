package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser() map[string][]string {
	lines := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	res := make(map[string][]string)
	for _, l := range lines {
		p := strings.Split(l, ":")
		o := strings.Fields(p[1])
		res[p[0]] = o
	}
	return res

}

var seen map[string]int

func p1_rec(g map[string][]string, act string) int {
	if act == "out" {
		return 1
	}
	mem, ok := seen[act]
	if ok {
		return mem
	}
	seen[act] = 0
	res := 0
	for _, c := range g[act] {
		res += p1_rec(g, c)
	}
	seen[act] = res
	return res
}

func part1(g map[string][]string) int {
	seen = make(map[string]int)
	return p1_rec(g, "you")
}

type seen_cell struct {
	s        string
	seen_dac bool
	seen_fft bool
}

var seen2 map[seen_cell]int

func p2_rec(g map[string][]string, act string, dac bool, fft bool) int {
	if act == "out" {
		if dac && fft {
			return 1
		}
		return 0
	}
	mem, ok := seen2[seen_cell{act, dac, fft}]
	if ok {
		return mem
	}
	res := 0
	dac = dac || (act == "dac")
	fft = fft || (act == "fft")

	seen2[seen_cell{act, dac, fft}] = 0
	for _, c := range g[act] {

		res += p2_rec(g, c, dac, fft)
	}
	seen2[seen_cell{act, dac, fft}] = res
	return res
}

func part2(g map[string][]string) int {
	seen2 = make(map[seen_cell]int)
	return p2_rec(g, "svr", false, false)
}

func main() {
	//fmt.Println(parser())
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
