package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed hyper_hyper_input_day_9.txt
var inputDay string

func parser() []int {
	var disk = strings.Split(inputDay, "")
	disk = disk[:len(disk)-1]
	var res []int = make([]int, 0)
	for i := range disk {
		n, _ := strconv.Atoi(disk[i])
		for range n {
			if i%2 == 0 {
				res = append(res, i/2)
			} else {
				res = append(res, -1)
			}
		}
	}
	return res
}

func checks(disk []int) int {
	var res int = 0
	for i := range disk {
		res += i * max(0, disk[i])
	}
	return res
}

func Part1(disk []int) int {
	var back_index int = len(disk) - 1
	var i int = 0
	for i < back_index {
		if disk[i] != -1 {
			i++
		} else {
			for disk[back_index] == -1 {
				back_index--
			}
			disk[i] = disk[back_index]
			disk[back_index] = -1
			back_index--
		}
	}
	return checks(disk)
}
func calc_size_block(disk []int, index int) int {
	val := disk[index]
	res := 0
	for disk[index] == val {
		res++
		index--
		if index < 0 {
			break
		}
	}
	return res
}
func calc_size_empty(disk []int, index int) int {
	res := 0
	for disk[index] == -1 {
		res++
		index++
	}
	return res
}
func Part2(disk []int) int {
	var back_index int = len(disk) - 1
	seen := make(map[int]bool)
	first_empty := 0
	for back_index >= 0 {
		if disk[back_index] == -1 || seen[disk[back_index]] {
			back_index--
		} else {
			seen[disk[back_index]] = true
			var i int = first_empty
			for disk[i] != -1 {
				i++
			}
			first_empty = i
            size_num := calc_size_block(disk, back_index)
			for i < back_index {
				size_empty := calc_size_empty(disk, i)
				if size_empty >= size_num {
					for range size_num {
						disk[i] = disk[back_index]
						disk[back_index] = -1
						i++
						back_index--
					}
					break
				} else {
					i += size_empty
					for disk[i] != -1 {
						i++
					}
				}
			}
			if i >= back_index {
				for back_index >= 0 && disk[back_index] != -1 && !seen[disk[back_index]] {
					back_index--

				}
			}
		}
	}
	return checks(disk)
}
func main() {
	fmt.Println("PART2:", Part2(parser()))
}
