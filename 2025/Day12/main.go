package main

import (
	_ "embed"
	"fmt"
	"strconv"

	//"slices"
	"strings"
)

//go:embed input2.txt
var inputDay string

func parser() ([][][]byte, [][2]int, [][]int) {
	all_filds := strings.Split(strings.TrimRight(inputDay, "\n"), "\n\n")
	shapes := make([][][]byte, 0)
	for _, shape := range all_filds[:len(all_filds)-1] {
		sh := make([][]byte, 0)
		lines := strings.Split(shape, "\n")
		for _, v := range lines[1:] {
			v2 := strings.Split(v, "")
			l := make([]byte, 0)
			for _, v3 := range v2 {
				l = append(l, v3[0])
			}
			sh = append(sh, l)
		}
		shapes = append(shapes, sh)
	}

	regions := strings.Split(all_filds[len(all_filds)-1], "\n")
	res_r := make([][]int, 0)
	size_res := make([][2]int, 0)
	for _, r := range regions {
		f := strings.Split(r, ": ")
		size := strings.Split(f[0], "x")
		i1, _ := strconv.Atoi(size[0])
		i2, _ := strconv.Atoi(size[1])
		size_res = append(size_res, [2]int{i1, i2})
		f2 := strings.Split(f[1], " ")
		res4 := make([]int, 0)
		for _, v := range f2 {
			i3, _ := strconv.Atoi(v)
			res4 = append(res4, i3)

		}
		res_r = append(res_r, res4)
	}
	return shapes, size_res, res_r

}

const (
	DEG0     = 0
	DEG90    = 1
	DEG180   = 2
	DEG270   = 3
	F_DEG0   = 4
	F_DEG90  = 5
	F_DEG180 = 6
	F_DEG270 = 7
)

type place_t struct {
	shape    int
	p        pos_t
	rotation int
}

type pos_t struct {
	x int
	y int
}

func apply_rota(p pos_t, rota int) pos_t {
	nx := p.x
	ny := p.y
	if rota > 3 {
		nx = -nx
		rota -= 4
	}
	for range rota {
		tmpx := nx
		tmpy := ny
		nx = tmpy
		ny = -tmpx
	}
	return pos_t{nx, ny}

}

func parse_shapes(sh [][][]byte) [][]pos_t {
	res := make([][]pos_t, 0)
	for _, s := range sh {
		shape := make([]pos_t, 0)
		for i, vi := range s {
			for j, vj := range vi {
				if vj == '#' {
					shape = append(shape, pos_t{i - 1, j - 1})
				}
			}
		}
		res = append(res, shape)
	}
	return res
}

func check_place(grid []place_t, newp place_t, size [2]int, sh [][]pos_t) bool {
	//fmt.Println("==============", grid, newp)
	for _, p := range sh[newp.shape] {
		p_rota := apply_rota(p, newp.rotation)
		abs_cco := pos_t{newp.p.x + p_rota.x, newp.p.y + p_rota.y}
		if abs_cco.x < 0 || abs_cco.y < 0 {
			return false
		}
		//fmt.Println("CHECK:", size, abs_cco)
		if abs_cco.x >= size[0] || abs_cco.y >= size[1] {
			return false
		}

		for _, other := range grid {
			if abs(other.p.x-newp.p.x) > 5 {
				continue
			}
			if abs(other.p.y-newp.p.y) > 5 {
				continue
			}
			for _, other_p := range sh[other.shape] {
				oter_rota := apply_rota(other_p, other.rotation)
				abs_other := pos_t{oter_rota.x + other.p.x, oter_rota.y + other.p.y}
				if abs_cco == abs_other {
					return false
				}
			}
		}
	}
	return true
}

func p1_rec(grid []place_t, x int, y int, all_sh [][]pos_t, size [2]int, regions []int, res int) int {
	//fmt.Println(x, y, res, grid)
	if x == size[0]-1 && y == size[1]-1 {
		//fmt.Println(res, grid)
		return res
	}
	next_cco := pos_t{x + ((y + 1) / size[1]), (y + 1) % size[1]}
	new_res := 0
	sum_re := 0
	for _, v := range regions {
		sum_re += v
	}
	for i, sh := range regions {
		if sh == 0 {
			continue
		}
		for rota := range 8 {
			new_sh := place_t{i, pos_t{x, y}, rota}
			if check_place(grid, new_sh, size, all_sh) {
				regions[i]--
				grid = append(grid, new_sh)
				//fmt.Println("Placed", i, x, y, rota)
				new_res = max(new_res, p1_rec(grid, next_cco.x, next_cco.y, all_sh, size, regions, res+1))
				if new_res >= sum_re {
					return new_res
				}
				regions[i]++
				grid = grid[:len(grid)-1]
			}
		}
	}
	return max(new_res, p1_rec(grid, next_cco.x, next_cco.y, all_sh, size, regions, res))
}

func part1(all_sh [][][]byte, sizes [][2]int, regions [][]int) int {
	better_sh := parse_shapes(all_sh)
	res := 0
	for i := range regions {
		//fmt.Println("===============================")
		//
		goal := 0
		for _, v := range regions[i] {
			goal += v
		}
		res_tmp := p1_rec(make([]place_t, 0), 0, 0, better_sh, sizes[i], regions[i], 0)
		fmt.Println(i)
		if res_tmp == goal {
			res += 1
		}
	}
	return res
}

func part2() int {
	return 0
}

func main() {
	fmt.Println(parser())
	//fmt.Println("Part1:", part1(parser()))
	//fmt.Println("Part2:", part2())
}
