package main

import (
    "fmt"
    _ "embed"
)

//go:embed input.txt
var inputDay string

func parser()*graph_t{
    grid := InToGrid(inputDay)
    var res *graph_t = nil
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] == "S"{
                res = rec_create_graph(grid,coo_t{i,j,RIGHT},0,make(map[node_id]*graph_t))
            }
        }
    }
    return res
}



const (
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3
    CLOCKWISE = 0
    COUNTERCLOCK = 1
)

type graph_t struct{
    coo_t
    cost int
    nei []*graph_t
}

type coo_t struct{
    x int
    y int
    dir int
}

var EMPTY graph_t= graph_t{coo_t{-1,-1,-1},-1,[]*graph_t{}}
var GOAL graph_t = graph_t{coo_t{-2,-2,-2},1,[]*graph_t{}}

func apply_dir(coo coo_t)coo_t{
    switch coo.dir{
    case RIGHT:
        return coo_t{coo.x,coo.y+1,coo.dir}
    case DOWN:
        return coo_t{coo.x+1,coo.y,coo.dir}
    case LEFT:
        return coo_t{coo.x,coo.y-1,coo.dir}
    case UP:
        return coo_t{coo.x-1,coo.y,coo.dir}
    default:
        return coo_t{-1,-1,-1}
    }
}

func turn(coo coo_t,cl int)coo_t{
    new_dir := 0
    if cl == CLOCKWISE{
        new_dir = (coo.dir +1) % 4
    }else{
        new_dir = coo.dir -1
        if new_dir == -1{
            new_dir = 3
        }
    }
    return coo_t{coo.x,coo.y,new_dir}
}

type node_id struct{
    coo_t
    cost int
}

func rec_create_graph(grid [][]string,coo coo_t,cost int,seen map[node_id]*graph_t)*graph_t{
    //fmt.Println(coo,cost)
    curr_node := node_id{coo,cost}
    if grid[coo.x][coo.y] == "#"{
        return &EMPTY
    }
    if seen[curr_node] != nil{
        return seen[curr_node]
    }
    if grid[coo.x][coo.y] == "E"{
        return &GOAL
    }
    res := graph_t{coo,cost,make([]*graph_t,0)}
    seen[curr_node] = &res
    next_co := apply_dir(coo)
    next_graph := rec_create_graph(grid,next_co,1,seen)
    if next_graph.coo_t != EMPTY.coo_t{
        res.nei = append(res.nei, next_graph)
    }
    next_co = turn(coo,CLOCKWISE)
    next_graph = rec_create_graph(grid,next_co,1000,seen)
    if next_graph.coo_t != EMPTY.coo_t{
        res.nei = append(res.nei, next_graph)
    }
    next_co = turn(coo,COUNTERCLOCK)
    next_graph = rec_create_graph(grid,next_co,1000,seen)
    if next_graph.coo_t != EMPTY.coo_t{
        res.nei = append(res.nei, next_graph)
    }
    return &res
}

func parcours(node *graph_t,dst_from_center int,dists map[*graph_t]int,seen map[*graph_t]bool)int{
    //fmt.Println(node.coo_t,node.cost,dst_from_center)
    delete (dists,node)
    if node.coo_t == GOAL.coo_t{
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

var prevs map[*graph_t]([]*graph_t) = make(map[*graph_t][]*graph_t)

func parcours_v2(node *graph_t,dst_from_center int,dists map[*graph_t]int,seen map[*graph_t]bool)*graph_t{
    //fmt.Println(node.coo_t,node.cost,dst_from_center)
    delete (dists,node)
    if node.coo_t == GOAL.coo_t{
        return node
    }
    seen[node] = true
    for _,nei := range node.nei{
        if seen[nei]{
            continue
        }
        _,ok := dists[nei]
        if !ok{
            dists[nei] = dst_from_center+ nei.cost
            prevs[nei] = append(prevs[nei], node)
        }else{
            if dst_from_center+nei.cost == dists[nei]{
                prevs[nei] = append(prevs[nei], node)
            }
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
    return parcours_v2(min_nei,dists[min_nei],dists,seen)
}

func parcours_prev(root *graph_t,marks map[coo_t]bool){
    only_coo := coo_t{root.x,root.y,0}
    marks[only_coo] = true
    //fmt.Println(root)
    nei,ok := prevs[root]
    if !ok{
        return
    }
    for _,next := range nei{
        parcours_prev(next,marks)
    }
    return
}

func Part2(g *graph_t)int{
    dists := make(map[*graph_t]int)
    seen := make(map[*graph_t]bool)
    dists[g] = 0
    var end *graph_t = parcours_v2(g,0,dists,seen)
    //fmt.Println(dists)
    var marks map[coo_t]bool = make(map[coo_t]bool)
    parcours_prev(end,marks)
    return len(marks)

}

func main(){
    fmt.Println("PART1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
