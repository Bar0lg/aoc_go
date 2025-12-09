package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser() [][]byte {
	res := make([][]byte, 0)
	g := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	for _, l := range g {
		caras := strings.Split(l, "")
		line := make([]byte, 0)
		for _, c := range caras {
			line = append(line, c[0])
		}
		res = append(res, line)
	}
	return res
}

type pos_t struct {
	x int
	y int
}

var seen = make(map[pos_t]bool)

func parcours_rec(g [][]byte, p pos_t) int {
	if p.x < 0 || p.x >= len(g) {
		return 0
	}
	if p.y < 0 || p.y >= len(g[0]) {
		return 0
	}
	if seen[p] == true {
		return 0
	}
	seen[p] = true
	if g[p.x][p.y] != '.' {
		tmp := parcours_rec(g, pos_t{p.x + 1, p.y + 1})
		tmp += parcours_rec(g, pos_t{p.x + 1, p.y - 1})
		return 1 + tmp
	}
	return parcours_rec(g, pos_t{p.x + 1, p.y})
}

func part1(g [][]byte) int {
	S := 0
	for i, v := range g[0] {
		if v == 'S' {
			S = i
			break
		}
	}
	res := parcours_rec(g, pos_t{0, S})
	return res
}

var seen2 = make(map[pos_t]int)

func parcours_rec2(g [][]byte, p pos_t) int {
	if p.x < 0 || p.x >= len(g) {
		return 0
	}
	if p.y < 0 || p.y >= len(g[0]) {
		return 0
	}
	mem, ok := seen2[p]
	if ok == true {
		return mem
	}
	if g[p.x][p.y] != '.' {
		tmp := parcours_rec2(g, pos_t{p.x + 1, p.y + 1})
		tmp += parcours_rec2(g, pos_t{p.x + 1, p.y - 1})
		seen2[p] = 1 + tmp
		return 1 + tmp
	}
	res_mem := parcours_rec2(g, pos_t{p.x + 1, p.y})
	seen2[p] = res_mem
	return res_mem
}

func part2(g [][]byte) int {
	S := 0
	for i, v := range g[0] {
		if v == 'S' {
			S = i
			break
		}
	}
	res := parcours_rec2(g, pos_t{0, S})
	return res + 1
}

func main() {
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
