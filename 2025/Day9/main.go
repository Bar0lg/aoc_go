package main

import (
	_ "embed"
	"fmt"
	"strconv"

	//"slices"
	"strings"
)

//go:embed test.txt
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

func give_up(list []pos_t, p1 pos_t, p2 pos_t) bool {
	lu := pos_t{min(p1.x, p2.x) - 1, min(p1.y, p2.y) - 1}
	ru := pos_t{min(p1.x, p2.x) - 1, max(p1.y, p2.y) + 1}
	rd := pos_t{max(p1.x, p2.x) + 1, max(p1.y, p2.y) + 1}
	for i := range list[:len(list)-1] {
		p1 := list[i]
		p2 := list[i+1]
		if p1.x == p2.x { // Cas vhorizon
			if p1.x > lu.x && p1.x < rd.x {
				if p1.y > lu.y && p1.y < ru.y && (p2.y <= lu.y || p2.y >= ru.y) {
					return false
				}
				if p2.y > lu.y && p2.y < ru.y && (p1.y <= lu.y || p1.y >= ru.y) {
					return false
				}
			}

		}
		if p1.y == p2.y { // Cas vertical
			if p1.y > lu.y && p1.y < ru.y {
				if p1.x > lu.x && p1.x < rd.x && (p2.x <= lu.x || p2.x >= rd.x) {
					return false
				}
				if p2.x > lu.x && p2.x < rd.x && (p1.x <= lu.x || p1.x >= rd.x) {
					return false
				}
			}

		}
	}

	return true

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
			if (p1.y <= p.y && p2.y > p.y) || (p2.y <= p.y && p1.y > p.y) {
				up++
			}
		}
		if p1.y == p2.y && p1.y < p.y {
			if (p1.x <= p.x && p2.x > p.x) || (p2.x <= p.x && p1.x > p.x) {
				side++
			}
		}

	}
	p1 := list[0]
	p2 := list[len(list)-1]
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
		if (p1.y <= p.y && p2.y > p.y) || (p2.y <= p.y && p1.y > p.y) {
			up++
		}
	}
	if p1.y == p2.y && p1.y < p.y {
		if (p1.x <= p.x && p2.x > p.x) || (p2.x <= p.x && p1.x > p.x) {
			side++
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
			can := give_up(poss, v1, v2)
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
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
