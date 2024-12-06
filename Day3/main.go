package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var InputDay string

func part1() int {
	var res int = 0
	var mulcatcher = regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	var all_muls []string = mulcatcher.FindAllString(InputDay, -1)
	var numCatcher = regexp.MustCompile("[0-9]{1,3}")
	for _, v := range all_muls {
		var nums []string = numCatcher.FindAllString(v, -1)
		var n1, _ = strconv.Atoi(nums[0])
		var n2, _ = strconv.Atoi(nums[1])
		res += n1 * n2
	}
	return res
}
func part2() int {
	var res int = 0
	var mulcatcher = regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")
	var all_muls []string = mulcatcher.FindAllString(InputDay, -1)
	var numCatcher = regexp.MustCompile("[0-9]{1,3}")
	var can_mult bool = true
	for _, v := range all_muls {
		if v == "do()" {
			can_mult = true
			continue
		}
		if v == "don't()" {
			can_mult = false
			continue
		}
		if !can_mult {
			continue
		}
		var nums []string = numCatcher.FindAllString(v, -1)
		var n1, _ = strconv.Atoi(nums[0])
		var n2, _ = strconv.Atoi(nums[1])
		res += n1 * n2
	}
	return res
}
func main() {
	fmt.Println(part2())
}
