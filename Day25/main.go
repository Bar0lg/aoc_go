package main

import (
    "fmt"
    "strings"
    _ "embed"
)

//go:embed input.txt
var inputDay string

func grid_to_vals(g [][]string,key bool)[]int{
    res := make([]int,0)
    for j := range g[0]{
        size := -1
        if !key{
            for i := range g{
                if g[i][j] == "#"{
                    size++
                }

            }
        }else{
            for i:=len(g)-1;i>0;i--{
                if g[i][j] == "#"{
                    size++
                }
            }
        }
        res = append(res,size)

    }
    return res
}

func parser()([][]int,[][]int){
    schemes := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    keys := make([][]int,0)
    locks := make([][]int,0)
    for _,v := range schemes{
        grid := InToGrid(v)
        if grid[0][0] == "#"{
            locks = append(locks,grid_to_vals(grid,false))

        }else{
            keys = append(keys,grid_to_vals(grid,true))
        }
    }
    return keys,locks
}

const LIMIT = 5

func Part1(keys [][]int,locks [][]int)int{
    res := 0
    flag := true
    for _,key := range keys{
        for _,lock := range locks{
            flag = true
            for pin := range key{
                if key[pin] + lock[pin] > LIMIT{
                    flag = false
                }
            }
            if flag{
            res++
            }
            
        }
    }
    return res
}
func Part2(){
    return
}

func main(){
    fmt.Println(parser())
    fmt.Println("PART1:",Part1(parser()))
    //fmt.Println("PART 2:",Part2(parser()))
}
