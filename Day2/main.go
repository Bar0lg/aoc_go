package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var InputDay string

func parser() [][]int {
	var reps []string = strings.Split(InputDay, "\n")
	reps = reps[:len(reps)-1]
	var res [][]int = make([][]int, 0)
	var levels []string = make([]string, 0)
	for _, v := range reps {
		res = append(res, make([]int, 0))
		levels = strings.Split(v, " ")
		for _, lev := range levels {
			lev_int, _ := strconv.Atoi(lev)
			res[len(res)-1] = append(res[len(res)-1], lev_int)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(input [][]int) int {
	var res int

	for _, reps := range input {
		var safe bool = true
		if reps[0] == reps[1] {
			continue
		}
		var desc bool = ((reps[0] - reps[1]) < 0)
		for i := range reps[:len(reps)-1] {
			if reps[i] == reps[i+1] {
				safe = false
			}
			if desc {
				if reps[i]-reps[i+1] > 0 {
					safe = false
				}
			} else {
				if reps[i]-reps[i+1] < 0 {
					safe = false
				}
			}

			if abs(reps[i]-reps[i+1]) > 3 {
				safe = false
			}
		}
		if safe {
			res += 1
		}

	}
	return res
}
func part2(input [][]int) int {
	var res int

	for _, reps := range input {
		if part1(append(make([][]int, 0), reps)) == 1 {
			res += 1
		} else {
			var corrected []int = make([]int, 0)
			for deleted_i := range reps {
				corrected = make([]int, 0)
				for lvl_i := range reps {
					if deleted_i == lvl_i {
						continue
					}
					corrected = append(corrected, reps[lvl_i])
				}

				if part1(append(make([][]int, 0), corrected)) == 1 {
					res += 1
					break
				}

			}
		}
	}
	return res
}

func main() {
	fmt.Println(part2(parser()))
}
