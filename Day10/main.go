package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
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

func sum_n(n int) int {
	return (n * (n + 1)) / 2
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
func calc_num_block(disk []int, index int) int {
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
func not_in(l []int, e int) bool {
	for _, v := range l {
		if v == e {
			return false
		}
	}
	return true
}
func Part2(disk []int) int {
	var back_index int = len(disk) - 1
	seen := make([]int, 0)
	for back_index >= 0 {
		if disk[back_index] == -1 || !(not_in(seen, disk[back_index])) {
			back_index--
		} else {
			seen = append(seen, disk[back_index])
			var i int = 0
			for disk[i] != -1 {
				i++
			}
			for i < back_index {
				size_empty := calc_size_empty(disk, i)
				size_num := calc_num_block(disk, back_index)
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
				for back_index >= 0 && disk[back_index] != -1 && not_in(seen, disk[back_index]) {
					back_index--

				}
			}
		}
	}
	return checks(disk)

}
func main() {
	fmt.Println(Part2(parser()))
}
