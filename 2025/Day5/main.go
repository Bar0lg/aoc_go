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

func parser() ([][]int, map[int]bool) {
	double := strings.Split(strings.TrimRight(inputDay, "\n"), "\n\n")
	ranges := strings.Split(double[0], "\n")
	res_ranges := make([][]int, 0)
	res_other := make(map[int]bool)
	for _, v := range ranges {
		inter := make([]int, 0)
		bornes := strings.Split(v, "-")
		i1, _ := strconv.Atoi(bornes[0])
		i2, _ := strconv.Atoi(bornes[1])
		inter = append(inter, i1)
		inter = append(inter, i2)
		res_ranges = append(res_ranges, inter)
	}
	fish := strings.Split(double[1], "\n")
	for _, v := range fish {
		i3, _ := strconv.Atoi(v)
		res_other[i3] = true
	}
	return res_ranges, res_other

}

func part1(all_ranges [][]int, fi map[int]bool) int {
	res := 0
	for key := range fi {
		for _, r := range all_ranges {
			if key >= r[0] && key <= r[1] {
				res++
				break
			}
		}
	}
	return res
}

func return_empty(r []int) []int {
	if r[0] > r[1] {
		return []int{-1, -1}
	}
	return r
}

func new_range(r1 []int, r2 []int) []int {
	if r2[0] >= r1[0] && r2[0] <= r1[1] {
		return return_empty([]int{r1[1] + 1, r2[1]})
	}
	if r2[1] >= r1[0] && r2[1] <= r1[1] {
		return return_empty([]int{r2[0], r1[0] - 1})
	}
	return r2
}

func if_inside(r1 []int, r2 []int) bool {
	if r2[0] >= r1[0] && r2[0] <= r1[1] && r2[1] >= r1[0] && r2[1] <= r1[1] {
		return true
	}
	return false
}

func part2(all_ranges [][]int) int {
	index := 1
	res := all_ranges[0][1] - all_ranges[0][0] + 1
	for index != len(all_ranges) {
		//fmt.Println("BEGIN")
		r := all_ranges[index]
		for i := 0; i < index; i++ {
			//fmt.Println(r, all_ranges[i])
			if !if_inside(all_ranges[i], r) {

				if if_inside(r, all_ranges[i]) {

					//fmt.Println("PROC")
					r1 := []int{r[0], all_ranges[i][0] - 1}
					r2 := []int{all_ranges[i][1] + 1, r[1]}
					all_ranges = append(all_ranges, r2)
					r = r1
					all_ranges[index] = r
					continue
				}

				r = new_range(all_ranges[i], r)
			} else {
				r = []int{-1, -1}
			}
			if r[0] == -1 {
				break
			}
		}
		//fmt.Println("HOOOOO", r)
		if r[0] != -1 {
			res += r[1] - r[0] + 1
		}
		index++
	}
	return res
}

func main() {
	p1, p2 := parser()
	fmt.Println("Part1:", part1(p1, p2))
	fmt.Println("Part2:", part2(p1))
}
