package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func InToGrid(input string) [][]string {
	var l []string = strings.Split(input, "\n")
	l = l[:len(l)-1]
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
	}
	return res
}
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
func parser() [][]int {
	return to_int(InToGrid(inputDay))
}

type coo_t struct {
	x int
	y int
}

func rec_parcours(grid [][]int, coo coo_t, seen map[coo_t]bool) int {
	if seen[coo] {
		return 0
	}
	seen[coo] = true
	if grid[coo.x][coo.y] == 9 {
		return 1
	}
	var res int = 0
	var val_pos int = grid[coo.x][coo.y]
	if coo.x >= 1 && grid[coo.x-1][coo.y] == val_pos+1 {
		res += rec_parcours(grid, coo_t{coo.x - 1, coo.y}, seen)
	}
	if coo.x < (len(grid)-1) && grid[coo.x+1][coo.y] == val_pos+1 {
		res += rec_parcours(grid, coo_t{coo.x + 1, coo.y}, seen)
	}
	if coo.y >= 1 && grid[coo.x][coo.y-1] == val_pos+1 {
		res += rec_parcours(grid, coo_t{coo.x, coo.y - 1}, seen)
	}
	if coo.y < (len(grid[0])-1) && grid[coo.x][coo.y+1] == val_pos+1 {
		res += rec_parcours(grid, coo_t{coo.x, coo.y + 1}, seen)
	}
	return res
}

func Part1(grid [][]int) int {
	var res int = 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				var seen map[coo_t]bool = make(map[coo_t]bool)
				res += rec_parcours(grid, coo_t{i, j}, seen)
			}
		}
	}
	return res

}
func rec_parcourspart2(grid [][]int, coo coo_t) int {
	if grid[coo.x][coo.y] == 9 {
		return 1
	}
	var res int = 0
	var val_pos int = grid[coo.x][coo.y]
	if coo.x >= 1 && grid[coo.x-1][coo.y] == val_pos+1 {
		res += rec_parcourspart2(grid, coo_t{coo.x - 1, coo.y})
	}
	if coo.x < (len(grid)-1) && grid[coo.x+1][coo.y] == val_pos+1 {
		res += rec_parcourspart2(grid, coo_t{coo.x + 1, coo.y})
	}
	if coo.y >= 1 && grid[coo.x][coo.y-1] == val_pos+1 {
		res += rec_parcourspart2(grid, coo_t{coo.x, coo.y - 1})
	}
	if coo.y < (len(grid[0])-1) && grid[coo.x][coo.y+1] == val_pos+1 {
		res += rec_parcourspart2(grid, coo_t{coo.x, coo.y + 1})
	}
	return res
}
func Part2(grid [][]int) int {
	var res int = 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				res += rec_parcourspart2(grid, coo_t{i, j})
			}
		}
	}
	return res
}

func main() {
	fmt.Println("PART1:", Part1(parser()))
	fmt.Println("PART 2:", Part2(parser()))
}
