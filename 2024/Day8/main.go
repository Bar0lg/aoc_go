package main

import (
    "fmt"
    "strings"
    _ "embed"
)

const (MAX_ANTENNA = 100)

//go:embed input.txt
var inputDay string

type coo_t struct{
    X int
    Y int
}

func InToGrid(input string) [][]string {
	var l []string = strings.Split(input, "\n")
	l = l[:len(l)-1]
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
	}
	return res
}

func parser()[][]string{
    return InToGrid(inputDay)
}

func create_map_of_antenna(grid [][]string) map[string]([MAX_ANTENNA]coo_t){
    var res map[string]([MAX_ANTENNA]coo_t) = make(map[string][MAX_ANTENNA]coo_t)
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] != "."{
                tmp := res[grid[i][j]]
                for tab_index,v := range tmp{
                    if (v == coo_t{0,0}){
                        tmp[tab_index] = coo_t{i,j}
                        break
                    }
                }
                res[grid[i][j]] = tmp
            }
        }
    }
    return res
}

func is_inbound(coo coo_t,bott_rigt coo_t)bool {
    if coo.X < 0 || coo.X > bott_rigt.X{
        return false
    }
    if coo.Y < 0 || coo.Y > bott_rigt.Y{
        return false
    }
    return true
}

func dist(c1 coo_t,c2 coo_t)coo_t{
    return coo_t{c1.X - c2.X,c1.Y-c2.Y}
}

func place_anti(nodes [MAX_ANTENNA]coo_t,antis map[coo_t]bool,br coo_t){
    for i := range nodes{
        if (nodes[i] == coo_t{0,0}){
            break
        }
        for j := range nodes{
            if (nodes[j] == coo_t{0,0}){
                break
            }
            if i == j{
                continue
            }
            anti_dist := dist(nodes[i],nodes[j])
            anti_pos := coo_t{nodes[i].X + anti_dist.X,nodes[i].Y + anti_dist.Y}
            if is_inbound(anti_pos,br){
                antis[anti_pos] = true
            }
        }
    }
}

func part1(grid [][]string)int{
    all_antinodes := make(map[coo_t]bool,0)
    map_all_node := create_map_of_antenna(grid)
    br := coo_t{len(grid)-1,len(grid[0])-1}
    for _,values := range map_all_node{
        place_anti(values,all_antinodes,br)
    }
    return len(all_antinodes)
}
func place_anti_part2(nodes [MAX_ANTENNA]coo_t,antis map[coo_t]bool,br coo_t){
    for i := range nodes{
        if (nodes[i] == coo_t{0,0}){
            break
        }
        for j := range nodes{
            if (nodes[j] == coo_t{0,0}){
                break
            }
            if i == j{
                continue
            }
            anti_dist := dist(nodes[i],nodes[j])
            anti_pos := coo_t{nodes[i].X,nodes[i].Y}
            k := 0
            for is_inbound(anti_pos,br){
                antis[anti_pos] = true
                anti_pos = coo_t{nodes[i].X + anti_dist.X*k,nodes[i].Y + anti_dist.Y*k}
                k++

            }
        }
    }
}
func part2(grid [][]string)int{
    all_antinodes := make(map[coo_t]bool,0)
    map_all_node := create_map_of_antenna(grid)
    br := coo_t{len(grid)-1,len(grid[0])-1}
    for _,values := range map_all_node{
        place_anti_part2(values,all_antinodes,br)
    }
    return len(all_antinodes)
}
func main(){
    fmt.Println(part2(parser()))
}

