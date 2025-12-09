package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"os/exec"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser() ([]byte, [][]int) {
	allops := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	matrix := make([][]int, 0)
	ops := make([]byte, 0)

	for i, v := range allops {
		if i == len(allops)-1 {
			by := strings.Fields(v)
			for _, a := range by {
				ops = append(ops, a[0])
			}
			continue
		}

		nums := strings.Fields(v)
		mat_line := make([]int, 0)
		for _, n := range nums {
			nint, _ := strconv.Atoi(n)
			mat_line = append(mat_line, nint)
		}
		matrix = append(matrix, mat_line)

	}
	return ops, matrix

}

func add(a int, b int) int {
	return a + b
}
func mul(a int, b int) int {
	return a * b
}

func part1(ops []byte, matrix [][]int) int {
	res := 0
	f := add
	for i, v := range ops {
		tmp := 0
		if v == '+' {
			f = add
			tmp = 0
		} else {
			f = mul
			tmp = 1
		}
		for _, num := range matrix {
			tmp = f(tmp, num[i])
		}
		res += tmp

	}
	return res
}

func part2(ops []byte, _ [][]int) int {

	out, _ := exec.Command("bash", "-c", "seq \"$(head -n1 input.txt |wc -c|cut -d' ' -f1)\"| xargs -I % sh -c \"cat input.txt |head -n-1 |cut -b% |tr -d '\n'&& echo '' \"").Output()
	numbers := string(out)
	all_nums := strings.Split(strings.TrimRight(numbers, "\n"), "\n    \n")
	fmt.Println(all_nums[0])
	new_mat := make([][]int, 0)
	for _, v := range all_nums {
		num2 := strings.Split(v, "\n")
		tmp := make([]int, 0)
		for _, v2 := range num2 {
			v2 = strings.Replace(v2, " ", "", -1)
			v2int, _ := strconv.Atoi(v2)
			tmp = append(tmp, v2int)

		}
		new_mat = append(new_mat, tmp)
	}
	fmt.Println(new_mat)
	res := 0
	f := add
	for i, v := range new_mat {
		tmp := 0
		if ops[i] == '+' {
			f = add
			tmp = 0
		} else {
			f = mul
			tmp = 1
		}
		for _, num := range v {
			tmp = f(tmp, num)
		}
		res += tmp

	}
	return res
}

func main() {
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
