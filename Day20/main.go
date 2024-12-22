package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputDay string

func parser()[][]string{
    return InToGrid(inputDay)
}



const (
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3
    LIMIT = 100
)

type coo_t struct{
    x int
    y int
}
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


func parcours_base(grid [][]string,dfb int,coo coo_t,dists map[coo_t]int){
    _,ok := dists[coo]
    if ok{
        return
    }
    if grid[coo.x][coo.y] == "#"{
        return
    }
    dists[coo] = dfb
    if grid[coo.x][coo.y] == "E"{
        return
    }
    for i := range(4){
        new_coo := apply_dir(coo,i)
        parcours_base(grid,dfb+1,new_coo,dists)
    }
    return
}

func parcours_cheat(g [][]string,dists map[coo_t]int)map[int]int{
    cheat := make(map[int]int)
    for key := range dists{
        for dir := range 4{
            next := apply_dir(key,dir)
            if g[next.x][next.y] != "#"{
                continue
            }
            new_coo := apply_dir(apply_dir(key,dir),dir)
            _,ok := dists[new_coo]
            if !ok{
                continue
            }
            if dists[new_coo] > dists[key]{
                cheat[dists[new_coo] - dists[key]-2]++
            }
        }
    }
    return cheat
}

func Part1(g [][]string)int{
    beg := coo_t{}
    for i := range g{
        for j := range g[0]{
            if g[i][j] == "S"{
                beg = coo_t{i,j}
            }
        }
    }
    //PrintGrid(g)
    dists := make(map[coo_t]int)
    parcours_base(g,0,beg,dists)
    all_chats := parcours_cheat(g,dists)
    res := 0
    for key,val := range all_chats{
        if key >= LIMIT{
            res += val
        }

    }

    return res


}
func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}

func parcours_cheat2(dists map[coo_t]int)map[int]int{
    cheat := make(map[int]int)
    for key := range dists{
        for d_x := -20;d_x<21;d_x++{
            for d_y := -20;d_y<21;d_y++{
                dist_man := abs(d_y) + abs(d_x)
                if dist_man > 20{
                    continue
                }
                new_coo := coo_t{key.x + (d_x),key.y + (d_y)}
                _,ok := dists[new_coo]
                if !ok{
                    continue
                }
                if dists[new_coo] > dists[key]{
                    //fmt.Println(key,new_coo,d_x,d_y,dist_man)
                    //fmt.Println(dists[new_coo] - dists[key]-dist_man)
                    //fmt.Println(key,new_coo)
                    cheat[dists[new_coo] - dists[key]-dist_man]++
                }
            }
    }
    }
    return cheat
}

func Part2(g [][]string)int{
    beg := coo_t{}
    for i := range g{
        for j := range g[0]{
            if g[i][j] == "S"{
                beg = coo_t{i,j}
            }
        }
    }
    //PrintGrid(g)
    dists := make(map[coo_t]int)
    parcours_base(g,0,beg,dists)
    all_chats := parcours_cheat2(dists)
    res := 0
    for key,val := range all_chats{
        if key >= LIMIT{
            res += val
        }

    }
    //fmt.Println(all_chats)

    return res


}

func main(){
    fmt.Println("PART 1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
