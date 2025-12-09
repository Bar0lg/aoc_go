package main

import (
	_ "embed"
	"fmt"
	"strconv"

	//"slices"
	"strings"
)

//go:embed input.txt
var inputDay string

type pos_t struct {
	x int
	y int
}

func parser() []pos_t {
	allp := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	res := make([]pos_t, 0)
	for _, p := range allp {
		ints := strings.Split(p, ",")
		i1, _ := strconv.Atoi(ints[0])
		i2, _ := strconv.Atoi(ints[1])
		res = append(res, pos_t{i1, i2})
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func area(p1 pos_t, p2 pos_t) int {
	return abs(p1.x-p2.x+1) * abs(p1.y-p2.y+1)
}

func part1(poss []pos_t) int {
	res := 0
	for i1, v1 := range poss {
		for _, v2 := range poss[i1+1:] {
			a := area(v1, v2)
			if a > res {
				res = a
			}
		}
	}
	return res
}

const (
	NOTHING = -1
	UP      = 0
	DOWN    = 2
	LEFT    = 1
	RIGHT   = 3
)

func what_move(p1 pos_t, p2 pos_t) int {
	if p1.x == p2.x {
		if p1.y > p2.y {
			return LEFT
		}
		return RIGHT
	}
	if p1.x > p2.x {
		return DOWN
	}
	return UP
}

func can_be_done(list []pos_t, p1 int, p2 int, last_m int, broke bool) bool {
	if p1 == len(list) {
		p1 = 0
	}
	if p2 == p1 {
		return true
	}
	pint := p1 + 1
	if pint == len(list) {
		pint = 0
	}
	m := what_move(list[p1], list[pint])
	if broke && (last_m == UP || last_m == DOWN) && (m == LEFT || m == RIGHT) {
		return false
	}
	if broke && (last_m == LEFT || last_m == RIGHT) && (m == UP || m == DOWN) {
		return false
	}
	if (last_m == LEFT || last_m == RIGHT) && (m == UP || m == DOWN) {
		return can_be_done(list, p1+1, p2, m, true)
	}
	if (last_m == UP || last_m == DOWN) && (m == LEFT || m == RIGHT) {
		return can_be_done(list, p1+1, p2, m, true)
	}
	return can_be_done(list, p1+1, p2, m, broke)
}

func is_green(list []pos_t, p pos_t) bool {
	side := 0
	up := 0
	for i := range list[:len(list)-1] {
		p1 := list[i]
		p2 := list[i+1]
		if p1.x == p2.x && p1.x == p.x {
			if (p1.y <= p.y && p2.y >= p.y) || (p2.y <= p.y && p1.y >= p.y) {
				return true
			}
		}
		if p1.y == p2.y && p1.y == p.y {
			if (p1.x <= p.x && p2.x >= p.x) || (p2.x <= p.x && p1.x >= p.x) {
				return true
			}
		}
		if p1.x == p2.x && p1.x < p.x {
			if (p1.y <= p.y && p2.y >= p.y) || (p2.y <= p.y && p1.y >= p.y) {
				up++
			}
		}
		if p1.y == p2.y && p1.y < p.y {
			if (p1.x <= p.x && p2.x >= p.x) || (p2.x <= p.x && p1.x >= p.x) {
				side++
			}
		}

	}
	if side%2 == 1 && up%2 == 1 {
		return true
	}
	return false
}

func part2(poss []pos_t) int {
	res := 0
	for i1, v1 := range poss {
		for i2, v2 := range poss {
			if i1 == i2 {
				continue
			}
			can := true
			for xi := min(v1.x, v2.x); xi <= max(v1.x, v2.x); xi++ {
				can = can && is_green(poss, pos_t{xi, v1.y})
				can = can && is_green(poss, pos_t{xi, v2.y})
			}
			for yi := min(v1.y, v2.y); yi <= max(v1.y, v2.y); yi++ {
				can = can && is_green(poss, pos_t{v1.x, yi})
				can = can && is_green(poss, pos_t{v2.x, yi})
			}
			if can {
				a := area(v1, v2)
				if a > res {
					res = a
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(parser())
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
