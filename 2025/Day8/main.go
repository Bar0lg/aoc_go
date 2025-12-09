package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"

	//"slices"
	"sort"
	"strings"
)

//go:embed test.txt
var inputDay string

type pos_t struct {
	x int
	y int
	z int
}

func parser() []pos_t {
	res := make([]pos_t, 0)
	poss := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	for _, v := range poss {
		indi_pos := strings.Split(v, ",")
		i1, _ := strconv.Atoi(indi_pos[0])
		i2, _ := strconv.Atoi(indi_pos[1])
		i3, _ := strconv.Atoi(indi_pos[2])
		p := pos_t{i1, i2, i3}
		res = append(res, p)
	}
	return res
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func pow2(x int) int {
	return x * x
}

func dist(p1 pos_t, p2 pos_t) float64 {
	return math.Sqrt(float64(pow2(p1.x-p2.x) + pow2(p1.y-p2.y) + pow2(p1.z-p2.z)))
}

type co_t struct {
	p1   pos_t
	p2   pos_t
	dist float64
}

func part1(boxs []pos_t) int {
	all_dists := make([]co_t, 0)
	for i, v1 := range boxs {
		for _, v2 := range boxs[i+1:] {
			tmp := co_t{v1, v2, dist(v1, v2)}
			all_dists = append(all_dists, tmp)
		}
	}
	sort.Slice(all_dists, func(i, j int) bool {
		return all_dists[i].dist < all_dists[j].dist
	})

	for i, v := range all_dists {
		if i == 0 {
			continue
		}
		if v.dist == all_dists[i-1].dist {
			fmt.Println("COLLISION:", v)
		}
	}

	groups := make(map[pos_t]int)
	curr := 1
	done := 0
	last := co_t{}
	for _, v := range all_dists {
		if done == 10 {
			break
		}
		if groups[v.p1] == groups[v.p2] && groups[v.p1] != 0 {
			done++
			continue
		}
		done++
		if groups[v.p1] == 0 && groups[v.p2] == 0 {
			groups[v.p1] = curr
			groups[v.p2] = curr
			curr++
			continue
		}
		if groups[v.p1] == 0 {
			groups[v.p1] = groups[v.p2]
			last = v
			continue
		}
		if groups[v.p2] == 0 {
			groups[v.p2] = groups[v.p1]
			last = v
			continue
		}

		list_to_change := make([]pos_t, 0)
		last = v
		for key_gr, val_gr := range groups {
			if val_gr == groups[v.p1] {
				list_to_change = append(list_to_change, key_gr)
			}
		}
		for _, pos_to_ch := range list_to_change {
			groups[pos_to_ch] = groups[v.p2]
		}

	}
	for _, v := range boxs {
		if groups[v] == 0 {
			groups[v] = curr
			curr++
		}
	}
	all_gr := make([]int, curr-1)
	for _, v := range groups {
		all_gr[v-1]++
	}
	slices.Sort(all_gr)
	fmt.Println(all_gr)
	fmt.Println(last)
	fmt.Println(groups)

	return all_gr[len(all_gr)-1] * all_gr[len(all_gr)-2] * all_gr[len(all_gr)-3]
}

func part2() int {
	return 0
}

func main() {
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2())
}
