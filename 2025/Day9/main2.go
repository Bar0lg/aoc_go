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
	fmt.Println("========================================", p1, p2, "==========================")
	up := min(p1.x, p2.x)
	down := max(p1.x, p2.x)
	right := max(p1.y, p2.y)
	left := min(p1.y, p2.y)
	for i := range list[:len(list)-1] {
		p1 := list[i]
		p2 := list[i+1]

		fmt.Println(p1, p2)
		if p1.x == p2.x { // Cas vhorizon

			if p1.x > up && p1.x < down {

				if p1.y > left && p1.y < right && (p2.y < left || p2.y > right) {
					return false
				}

				if p2.y > left && p2.y < right && (p1.y < left || p1.y > right) {
					return false
				}
			}

		}
		if p1.y == p2.y { // Cas vhorizon

			if p1.y > left && p1.y < right {

				if p1.x > up && p1.x < down && (p2.x < up || p2.x > down) {
					return false
				}
				if p2.x > up && p2.x < down && (p1.x < up || p1.x > down) {
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
			fmt.Println(can)
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
