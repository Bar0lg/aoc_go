package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser()*graph_t{
    grid := make([][]string,0)
    for range MAX_X{
        tmp := make([]string,0)
        for range MAX_Y{
            tmp = append(tmp, ".")
        }
        grid = append(grid, tmp)
    }
    var res *graph_t = nil
    corr := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    for i := range 1024{
        corr_coo := strings.Split(corr[i],",")
        n1,_ := strconv.Atoi(corr_coo[0])
        n2,_ := strconv.Atoi(corr_coo[1])
        grid[n2][n1] = "#"
    }
    //PrintGrid(grid)
    res = rec_create_graph(grid,coo_t{0,0},0,make(map[coo_t]*graph_t))
    return res
}



const (
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3
    MAX_Y = 71
    MAX_X = 71
)

type graph_t struct{
    coo_t
    cost int
    nei []*graph_t
}

type coo_t struct{
    x int
    y int
}

var EMPTY graph_t= graph_t{coo_t{-1,-1},-1,[]*graph_t{}}

func apply_dir(coo coo_t,dir int)coo_t{
    switch dir{
    case RIGHT:
        return coo_t{coo.x,coo.y+1}
    case DOWN:
        return coo_t{coo.x+1,coo.y}
    case LEFT:
        return coo_t{coo.x,coo.y-1}
    case UP:
        return coo_t{coo.x-1,coo.y}
    default:
        return coo_t{-1,-1}
    }
}

func rec_create_graph(grid [][]string,coo coo_t,cost int,seen map[coo_t]*graph_t)*graph_t{
    //fmt.Println(coo,cost)
    curr_node := coo
    if coo.x <0 || coo.x > MAX_X -1{
        return &EMPTY
    }
    if coo.y < 0 || coo.y > MAX_Y -1{
        return &EMPTY
    }
    if grid[coo.x][coo.y] == "#"{
        return &EMPTY
    }
    if seen[curr_node] != nil{
        return seen[curr_node]
    }
    res := graph_t{coo,cost,make([]*graph_t,0)}
    seen[curr_node] = &res
    for i := range(4){
        next_co := apply_dir(coo,i)
        next_graph := rec_create_graph(grid,next_co,1,seen)
        if next_graph.coo_t != EMPTY.coo_t{
            res.nei = append(res.nei, next_graph)
        }
}
    return &res
}

func parcours(node *graph_t,dst_from_center int,dists map[*graph_t]int,seen map[*graph_t]bool)int{
    //fmt.Println(node.coo_t,node.cost,dst_from_center)
    delete (dists,node)
    if (node.coo_t == coo_t{MAX_X-1,MAX_Y-1}){
        return dst_from_center
    }
    seen[node] = true
    for _,nei := range node.nei{
        if seen[nei]{
            continue
        }
        _,ok := dists[nei]
        if !ok{
            dists[nei] = dst_from_center+ nei.cost
        }else{
            dists[nei] = min(dists[nei],dst_from_center+nei.cost)
        }
    }
    min_dst := -1
    var min_nei *graph_t = nil
    for key,val := range dists{
        if min_dst == -1{
            min_nei = key
            min_dst = val
        }
        if min_dst > val{
            min_dst = val
            min_nei = key
        }
    }
    //fmt.Println("MIN:",min_nei)
    return parcours(min_nei,dists[min_nei],dists,seen)
}

func Part1(g *graph_t)int{
    dists := make(map[*graph_t]int)
    seen := make(map[*graph_t]bool)
    dists[g] = 0
    var res int = parcours(g,0,dists,seen)
    //fmt.Println(dists)
    return res


}
func parcoursv2(grid [][]string,coo coo_t,seen map[coo_t]bool)bool{
    //fmt.Println(node.coo_t,node.cost,dst_from_center)
    if seen[coo]{
        return false
    }
    if coo.x < 0 || coo.x >= MAX_X{
        return false
    }
    if coo.y < 0 || coo.y >= MAX_Y{
        return false
    }
    if grid[coo.x][coo.y] == "#"{
        return false
    }
    if (coo == coo_t{MAX_X-1,MAX_Y-1}){
        return true
    }
    seen[coo] = true
    res := false
    for i := range 4{
        new_coo := apply_dir(coo,i)
        res = res || parcoursv2(grid,new_coo,seen)
    }
    return res
}

func Part2(){
    grid := make([][]string,0)
    for range MAX_X{
        tmp := make([]string,0)
        for range MAX_Y{
            tmp = append(tmp, ".")
        }
        grid = append(grid, tmp)
    }
    corr := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    for i := range corr{
        corr_coo := strings.Split(corr[i],",")
        n1,_ := strconv.Atoi(corr_coo[0])
        n2,_ := strconv.Atoi(corr_coo[1])
        grid[n2][n1] = "#"
        if !parcoursv2(grid,coo_t{0,0},make(map[coo_t]bool)){
            fmt.Printf("PART2:%d,%d\n",n1,n2)
            return
        }
    }
    //PrintGrid(grid)
    return

}

func main(){
    fmt.Println("PART1:",Part1(parser()))
    Part2()
}
