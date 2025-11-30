package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
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

func parser() [][]string {
	return InToGrid(inputDay)
}

type coo_t struct {
	x int
	y int
}
type Empty struct{}

func rec_area(grid [][]string, region string, coo coo_t, seen map[coo_t]Empty) int {
	if grid[coo.x][coo.y] != region {
		return 0
	}
	_, ok := seen[coo]
	if ok {
		return 0
	}
	var res int = 0
	seen[coo] = Empty{}
	if coo.x > 0 {
		res += rec_area(grid, region, coo_t{coo.x - 1, coo.y}, seen)
	}
	if coo.x < len(grid)-1 {
		res += rec_area(grid, region, coo_t{coo.x + 1, coo.y}, seen)
	}
	if coo.y > 0 {
		res += rec_area(grid, region, coo_t{coo.x, coo.y - 1}, seen)
	}
	if coo.y < len(grid[0])-1 {
		res += rec_area(grid, region, coo_t{coo.x, coo.y + 1}, seen)
	}
	return 1 + res
}
func rec_perimeter(grid [][]string, region string, coo coo_t, seen map[coo_t]Empty) int {
	if grid[coo.x][coo.y] != region {
		return 1
	}
	_, ok := seen[coo]
	if ok {
		return 0
	}
	seen[coo] = Empty{}
	var res int = 0
	if coo.x > 0 {
		res += rec_perimeter(grid, region, coo_t{coo.x - 1, coo.y}, seen)
	} else {
		res += 1
	}
	if coo.x < len(grid)-1 {
		res += rec_perimeter(grid, region, coo_t{coo.x + 1, coo.y}, seen)
	} else {
		res += 1
	}
	if coo.y > 0 {
		res += rec_perimeter(grid, region, coo_t{coo.x, coo.y - 1}, seen)
	} else {
		res += 1
	}
	if coo.y < len(grid[0])-1 {
		res += rec_perimeter(grid, region, coo_t{coo.x, coo.y + 1}, seen)
	} else {
		res += 1
	}
	return res
}
func Part1(grid [][]string) int {
	res := 0
	seen := make(map[coo_t]Empty)
	seen_peri := make(map[coo_t]Empty)

	for i := range grid {
		for j := range grid[i] {
            _,ok := seen[coo_t{i,j}]
            if !ok{
                area := rec_area(grid, grid[i][j], coo_t{i, j}, seen)
				perimter := rec_perimeter(grid, grid[i][j], coo_t{i, j}, seen_peri)
				res += area * perimter
			}
		}
	}
	return res
}

const (
	DROITE = 0
	BAS    = 1
	GAUCHE = 2
	HAUT   = 3
)

func convex_angle(grid [][]string, coo coo_t, region string) int {
	res := 0
	if coo.x == 0 || grid[coo.x-1][coo.y] != region {
		if coo.y == len(grid[0])-1 || grid[coo.x][coo.y+1] != region {
			res += 2
		}
	}
	if coo.x == len(grid)-1 || grid[coo.x+1][coo.y] != region {
		if coo.y == 0 || grid[coo.x][coo.y-1] != region {
			res += 2
		}
	}
	return res
}

func concave_angles(grid [][]string, coo coo_t, region string) int {
	res := 0
	if coo.x > 0 && grid[coo.x-1][coo.y] == region {
		if coo.y > 0 && grid[coo.x][coo.y-1] == region {
			if grid[coo.x-1][coo.y-1] != region {
				res += 2
			}
		}
	}
	if coo.x < len(grid)-1 && grid[coo.x+1][coo.y] == region {
		if coo.y < len(grid)-1 && grid[coo.x][coo.y+1] == region {
			if grid[coo.x+1][coo.y+1] != region {
				res += 2
			}
		}
	}
	return res
}

func rec_sides(grid [][]string, region string, coo coo_t, seen map[coo_t]Empty) int {
	if grid[coo.x][coo.y] != region {
		return 0
	}
	_, ok := seen[coo]
	if ok {
		return 0
	}
	seen[coo] = Empty{}
	var res int = 0
	if coo.x > 0 {
		res += rec_sides(grid, region, coo_t{coo.x - 1, coo.y}, seen)
	}
	if coo.x < len(grid)-1 {
		res += rec_sides(grid, region, coo_t{coo.x + 1, coo.y}, seen)
	}
	if coo.y > 0 {
		res += rec_sides(grid, region, coo_t{coo.x, coo.y - 1}, seen)
	}
	if coo.y < len(grid[0])-1 {
		res += rec_sides(grid, region, coo_t{coo.x, coo.y + 1}, seen)
	}
	return res + convex_angle(grid, coo, region) + concave_angles(grid, coo, region)
}
func Part2(grid [][]string) int {
	res := 0
	seen := make(map[coo_t]Empty)
	seen_angles := make(map[coo_t]Empty)

	for i := range grid {
		for j := range grid[i] {
            _,ok := seen[coo_t{i,j}]
            if !ok{
                area := rec_area(grid, grid[i][j], coo_t{i, j}, seen)
                sides := rec_sides(grid, grid[i][j], coo_t{i, j}, seen_angles)
                res += area * sides
			}
		}
	}
	return res
}

func main() {
	//fmt.Println(parser())
	fmt.Println("PART1:", Part1(parser()))
	fmt.Println("PART 2:", Part2(parser()))
}
