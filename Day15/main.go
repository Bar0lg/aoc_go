package main

import (
    "fmt"
    "strings"
    _ "embed"
)

type coo_t struct{
    x int
    y int
}
const (
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3
)

//go:embed input.txt
var inputDay string


func InToGrid(input string) [][]string {
	var l []string = strings.Split(input, "\n")
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
	}
	return res
}

func parser()([][]string,[]string){
    map_and_movs := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    grid := map_and_movs[0]
    movs := map_and_movs[1]
    var movs_continued []string = make([]string, 0)
    for _,v := range strings.Split(movs,""){
        if v != "\n"{
            movs_continued = append(movs_continued, v)
        }
    }
    return InToGrid(grid),movs_continued
}

func push(grid [][]string,coo coo_t,dir int)coo_t{
    var next coo_t = coo_t{}
    switch dir{
    case RIGHT:
        next = coo_t{coo.x,coo.y+1}
    case DOWN:
        next = coo_t{coo.x+1,coo.y}
    case LEFT:
        next = coo_t{coo.x,coo.y-1}
    case UP:
        next = coo_t{coo.x-1,coo.y}
    default:
        next = coo_t{-1,-1}
    }
    if grid[next.x][next.y] == "#"{
        return coo_t{-1,-1}
    }
    if grid[next.x][next.y] == "."{
        tmp := grid[coo.x][coo.y]
        grid[coo.x][coo.y] = grid[next.x][next.y]
        grid[next.x][next.y] = tmp
        return next
    }
    if grid[next.x][next.y] == "O"{
        status := push(grid,next,dir)
        if (status != coo_t{-1,-1}){
            tmp := grid[coo.x][coo.y]
            grid[coo.x][coo.y] = grid[next.x][next.y]
            grid[next.x][next.y] = tmp
            return next
        }
    }
    return coo_t{-1,-1}
}

func move(grid [][]string,coo coo_t,moves []string){
    if len(moves) == 0{
        return
    }
    //print_grid(grid)
    var next coo_t = coo_t{}
    var dir int = 0
    switch moves[0]{
    case ">":
        next = coo_t{coo.x,coo.y+1}
        dir = RIGHT
    case "v":
        next = coo_t{coo.x+1,coo.y}
        dir =DOWN
    case "<":
        next = coo_t{coo.x,coo.y-1}
        dir = LEFT
    case "^":
        next = coo_t{coo.x-1,coo.y}
        dir = UP
    default:
        next = coo_t{-1,-1}
    }
    if grid[next.x][next.y] == "#"{
        move(grid,coo,moves[1:])
        return
    }
    if grid[next.x][next.y] == "."{
        tmp := grid[coo.x][coo.y]
        grid[coo.x][coo.y] = grid[next.x][next.y]
        grid[next.x][next.y] = tmp
        move(grid,next,moves[1:])
        return
    }
    if grid[next.x][next.y] == "O"{
        status := push(grid,next,dir)
        if (status != coo_t{-1,-1}){
            tmp := grid[coo.x][coo.y]
            grid[coo.x][coo.y] = grid[next.x][next.y]
            grid[next.x][next.y] = tmp
            move(grid,next,moves[1:])
            return
        }
        move(grid,coo,moves[1:])
    }
    return
}

func print_grid(g [][]string){
    fmt.Printf("\n")
    for _,vi := range g{
        for _,vj := range vi{
            fmt.Printf("%s",vj)
        }
        fmt.Printf("\n")
    } 
}

func GPS(i,j int)int{
    return 100 *i + j
}

func Part1(grid [][]string,moves []string)int{
    char_pos := coo_t{}
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] == "@"{
                char_pos = coo_t{i,j}
            }
        }
    }
    move(grid,char_pos,moves)
    res := 0
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] == "O"{
                res += GPS(i,j)
            }
        }
    }
    return res
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

func push_v2(grid [][]string,coo coo_t,dir int)int{
    var next coo_t = apply_dir(coo,dir)
    
    if grid[next.x][next.y] == "#"{
        return -1
    }
    if grid[next.x][next.y] == "."{
        tmp := grid[coo.x][coo.y]
        grid[coo.x][coo.y] = grid[next.x][next.y]
        grid[next.x][next.y] = tmp
        return 0
    }
    if grid[next.x][next.y] == "[" || grid[next.x][next.y] == "]"{
        other := coo_t{-1,-1}
        if grid[next.x][next.y] == "["{
            other = apply_dir(next,RIGHT)
        }else{
            other = apply_dir(next,LEFT)
        }
        if dir == UP || dir == DOWN{
            s1 := push_v2(grid,next,dir)
            s2 := push_v2(grid,other,dir)
            if s1 == -1 && s2 == -1{
                return -1
            }
            if s1 == -1{
                return -1
            }
            if s2 == -1{
                return -1
            }
            tmp := grid[coo.x][coo.y]
            grid[coo.x][coo.y] = grid[next.x][next.y]
            grid[next.x][next.y] = tmp
            return 0
        }else{
            other = apply_dir(next,dir)
            status := push_v2(grid,other,dir)
            if status != -1{
                grid[next.x][next.y],grid[other.x][other.y] = grid[other.x][other.y],grid[next.x][next.y]
                grid[next.x][next.y],grid[coo.x][coo.y] = grid[coo.x][coo.y],grid[next.x][next.y]
                return 0
            }
            return -1
        }
    }
    fmt.Println(coo,"GROS PROBLEME")
    return -1
}

func move_v2(grid [][]string,coo coo_t,moves []string)[][]string{
    if len(moves) == 0{
        return grid
    }
    //print_grid(grid)
    var next coo_t = coo_t{}
    var dir int = 0
    switch moves[0]{
    case ">":
        next = coo_t{coo.x,coo.y+1}
        dir = RIGHT
    case "v":
        next = coo_t{coo.x+1,coo.y}
        dir =DOWN
    case "<":
        next = coo_t{coo.x,coo.y-1}
        dir = LEFT
    case "^":
        next = coo_t{coo.x-1,coo.y}
        dir = UP
    default:
        next = coo_t{-1,-1}
    }
    if grid[next.x][next.y] == "#"{
        return move_v2(grid,coo,moves[1:])
    }
    if grid[next.x][next.y] == "."{
        tmp := grid[coo.x][coo.y]
        grid[coo.x][coo.y] = grid[next.x][next.y]
        grid[next.x][next.y] = tmp
        return move_v2(grid,next,moves[1:])
    }
    if grid[next.x][next.y] == "[" || grid[next.x][next.y] == "]"{
        backup_grid := copy_grid(grid)
        other := coo_t{-1,-1}
        if grid[next.x][next.y] == "["{
            other = apply_dir(next,RIGHT)
        }else{
            other = apply_dir(next,LEFT)
        }
        if dir == UP || dir == DOWN{
            s1 := push_v2(grid,next,dir)
            s2 := push_v2(grid,other,dir)
            if s1 == 0 && s2 == 0{
                grid[coo.x][coo.y],grid[next.x][next.y] = grid[next.x][next.y],grid[coo.x][coo.y]
                return move_v2(grid,next,moves[1:])
            }
            return move_v2(backup_grid,coo,moves[1:])
        }else{
            other = apply_dir(next,dir)
            status := push_v2(grid,other,dir)
            if status != -1{
                grid[next.x][next.y],grid[other.x][other.y] = grid[other.x][other.y],grid[next.x][next.y]
                grid[next.x][next.y],grid[coo.x][coo.y] = grid[coo.x][coo.y],grid[next.x][next.y]
                return move_v2(grid,next,moves[1:])
            }
            return move_v2(backup_grid,coo,moves[1:])
        }
    }
    fmt.Println(coo,"GROS PROBLEME")
    return make([][]string,0)
}

func grid_p2(g [][]string)[][]string{
    res := make([][]string,0)
    for _,vi := range g{
        tmp := make([]string,0)
        for _,vj := range vi{
            switch vj{
            case "#":
                tmp = append(tmp, "#")
                tmp = append(tmp, "#")
            case ".":
                tmp = append(tmp, ".")
                tmp = append(tmp, ".")
            case "@":
                tmp = append(tmp, "@")
                tmp = append(tmp, ".")
            case "O":
                tmp = append(tmp, "[")
                tmp = append(tmp, "]")
            }
        }
        res = append(res, tmp)
    }
    return res
}

func copy_grid(g [][]string)[][]string{
    res := make([][]string,0)
    for _,vi := range g{
        tmp := make([]string,0)
        for _,vj := range vi{
            tmp = append(tmp, vj)
        }
        res = append(res, tmp)
    }
    return res
}

func Part2(grid [][]string,moves []string)int{
    char_pos := coo_t{}
    grid = grid_p2(grid)
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] == "@"{
                char_pos = coo_t{i,j}
            }
        }
    }
    grid = move_v2(grid,char_pos,moves)
    res := 0
    for i := range grid{
        for j := range grid[i]{
            if grid[i][j] == "[" {
                res += GPS(i,j)
            }
        }
    }
    return res
}
func main(){
    //fmt.Println(parser())
    fmt.Println("PART1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
