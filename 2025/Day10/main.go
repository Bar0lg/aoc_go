package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"

	//"slices"
	"github.com/mitchellh/go-z3"
	"strings"
)

//go:embed input.txt
var inputDay string

type machine_t struct {
	lights  []int
	buttons [][]int
	jolt    []int
}

func parser() []machine_t {
	mech_res := make([]machine_t, 0)
	all_mech := strings.Split(strings.TrimRight(inputDay, "\n"), "\n")
	for _, m := range all_mech {
		fil := strings.Fields(m)

		li := strings.Split(fil[0][1:len(fil[0])-1], "")
		li_res := make([]int, 0)
		for i, l := range li {
			if l[0] == '#' {
				li_res = append(li_res, i)
			}
		}
		all_buts := make([][]int, 0)
		for _, butt := range fil[1 : len(fil)-1] {
			butt_spl := strings.Split(butt[1:len(butt)-1], ",")
			butt_res := make([]int, 0)
			for _, v := range butt_spl {
				a, _ := strconv.Atoi(v)
				butt_res = append(butt_res, a)
			}
			all_buts = append(all_buts, butt_res)
		}
		jo := strings.Split(fil[len(fil)-1][1:len(fil[len(fil)-1])-1], ",")
		jo_res := make([]int, 0)
		for _, l := range jo {
			a, _ := strconv.Atoi(l)
			jo_res = append(jo_res, a)
		}

		mech_res = append(mech_res, machine_t{li_res, all_buts, jo_res})
	}
	return mech_res

}

func light_to_int(lights []int) int {
	res := 0
	for _, l := range lights {
		res = res | (1 << l)
	}
	return res
}

func int_to_l(l int) []int {
	res := make([]int, 0)
	i := 0
	for l != 0 {
		if l%2 == 1 {
			res = append(res, i)
		}
		l = l >> 1
		i++
	}
	return res
}

func apply_button(l int, but []int) int {
	li := int_to_l(l)
	for _, v := range but {
		if slices.Contains(li, v) {
			i := slices.Index(li, v)
			li = slices.Delete(li, i, i+1)
		} else {
			li = append(li, v)
		}
	}
	return light_to_int(li)
}

func bfs1(ma machine_t) map[int]int {
	seen := make(map[int]bool)
	to_see := make([]int, 0)
	parents := make(map[int]int)
	found := false
	queue_index := 0

	to_see = append(to_see, 0)

	goal_int := light_to_int(ma.lights)

	for !found {
		state := to_see[queue_index]
		queue_index++
		seen[state] = true
		if state == goal_int {
			found = true
			continue
		}
		for _, buts := range ma.buttons {
			nei := apply_button(state, buts)
			if !seen[nei] {
				parents[nei] = state
				seen[nei] = true
				to_see = append(to_see, nei)
			}
		}

	}
	return parents
}

func length_parcours(parents map[int]int, goal int) int {
	res := 0
	s := goal
	for s != 0 {
		s = parents[s]
		res++
	}
	return res
}

func part1(machs []machine_t) int {
	res := 0
	for _, m := range machs {
		res += length_parcours(bfs1(m), light_to_int(m.lights))
	}
	return res
}

func volt_to_str(jo []int) string {
	arr := make([]string, 0)
	for _, v := range jo {
		arr = append(arr, strconv.Itoa(v))
	}
	return strings.Join(arr, "-")
}

func str_to_jo(s string) []int {
	res := make([]int, 0)
	spl := strings.Split(s, "-")
	for _, v := range spl {
		a, _ := strconv.Atoi(v)
		res = append(res, a)
	}
	return res
}

func apply_jo_buttons(js string, bu []int) string {
	jo := str_to_jo(js)
	for _, v := range bu {
		jo[v]++
	}
	return volt_to_str(jo)
}

func h(jo string, g []int) int {
	jo_arr := str_to_jo(jo)
	res := jo_arr[0]
	for i, v := range jo_arr {
		hi := g[i] - v
		if hi < 0 {
			return 9999999
		}
		res = max(hi, res)
	}
	return res
}

func bfs2(ma machine_t) map[string]string {
	seen := make(map[string]bool)
	to_see := make(map[string]int)
	parents := make(map[string]string)
	found := false
	queue_index := 0
	beg := volt_to_str(make([]int, len(ma.jolt)))
	to_see[beg] = h(beg, ma.jolt)

	goal_int := volt_to_str(ma.jolt)

	//fmt.Println(goal_int, "->", beg)

	for !found {
		state := ""
		h_st := -1
		for k, v := range to_see {
			if h_st == -1 || v < h_st {
				state = k
				h_st = v
			}
		}

		dist := length_parcoursp2(parents, state)
		queue_index++
		seen[state] = true
		delete(to_see, state)
		//fmt.Println(state)
		if state == goal_int {
			found = true
			continue
		}
		for _, buts := range ma.buttons {
			nei := apply_jo_buttons(state, buts)
			if !seen[nei] {
				parents[nei] = state
				seen[nei] = true
				to_see[nei] = dist + h(nei, ma.jolt)
			}
		}

	}
	fmt.Println("HIII")
	return parents
}

func length_parcoursp2(parents map[string]string, goal string) int {
	res := 0
	s := goal
	beg := volt_to_str(make([]int, len(str_to_jo(goal))))
	for s != beg {
		s = parents[s]
		res++
	}
	return res
}

func part2(machs []machine_t) int {
	res := 0
	for _, m := range machs {
		res += length_parcoursp2(bfs2(m), volt_to_str(m.jolt))
	}
	return res
}

func main() {
	//fmt.Println(parser())
	fmt.Println("Part1:", part1(parser()))
	fmt.Println("Part2:", part2(parser()))
}
