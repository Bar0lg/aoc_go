package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var inputDay string

func parser()[][]string{
    return InToGrid(inputDay)
}



const (
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3
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
    if dists[coo] != 0{
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
    for key,_ := range dists{
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
                cheat[dists[new_coo] - dists[key]]++
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
    dists := make(map[coo_t]int)
    parcours_base(g,0,beg,dists)
    fmt.Println(parcours_cheat(g,dists))

    return 0


}

func main(){
    fmt.Println("PART1:",Part1(parser()))
    //Part2()
}
