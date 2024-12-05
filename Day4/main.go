package main

import(
    "fmt"
    "strings"
    _ "embed"
)

//go:embed input.txt
var InputDay string

type coo_t struct{
    X int
    Y int
}

var R = coo_t{0,1}
var L = coo_t{0,-1}
var U = coo_t{-1,0}
var D = coo_t{1,0}
var UR = coo_t{-1,1}
var UL = coo_t{-1,-1}
var DR = coo_t{1,1}
var DL = coo_t{1,-1}

var dirs [8]coo_t = [8]coo_t{R,L,U,D,UR,UL,DR,DL}

func apply_dir(coo coo_t,dir coo_t, max_line int,max_col int) coo_t{
    if coo.X + dir.X >= max_line || coo.X + dir.X <0{
        return coo_t{-1,-1}
    }
    if coo.Y + dir.Y >= max_col || coo.Y + dir.Y <0{
        return coo_t{-1,-1}
    }
    return coo_t{coo.X+dir.X,coo.Y+dir.Y}
}

func parser()[][]string{
    var lines []string = strings.Split(InputDay,"\n")
    lines = lines[:len(lines)-1]
    var res [][]string = make([][]string,0)
    for _,v := range lines{
        var letters []string = strings.Split(v,"")
        res = append(res, letters)
    }
    return res
}

var WORD [4]string = [4]string{"X","M","A","S"}

func go_catch(tab [][]string,word_index int,coo coo_t,dir coo_t) bool{
    if (coo == coo_t{-1,-1}){
        return false
    }
    if (word_index == 3 && tab[coo.X][coo.Y] == WORD[word_index]){
        return true
    }
    if tab[coo.X][coo.Y] != WORD[word_index]{
        return false
    }
    return go_catch(tab,word_index+1,apply_dir(coo,dir,len(tab),len(tab[0])),dir)

}

func part1(t [][]string)int{
    var res int =0
    for i := range t{
        for j := range t[i]{
            if t[i][j] != WORD[0]{
                continue
            }
            for _,d := range dirs{
                if go_catch(t,0,coo_t{i,j},d){
                    res += 1
                }   
            }
        }
    }
    return res
}
var coo_oob coo_t = coo_t{-1,-1}

func part2(t [][]string)int{
    var res int =0
    for i := range t{
        for j := range t[i]{
            if t[i][j] != "A"{
                continue
            }
            var UL_coo coo_t = apply_dir(coo_t{i,j},UL,len(t),len(t[0]))
            var UR_coo coo_t = apply_dir(coo_t{i,j},UR,len(t),len(t[0]))
            var DL_coo coo_t = apply_dir(coo_t{i,j},DL,len(t),len(t[0]))
            var DR_coo coo_t = apply_dir(coo_t{i,j},DR,len(t),len(t[0]))
            if UL_coo == coo_oob || UR_coo == coo_oob || DL_coo == coo_oob || DR_coo == coo_oob{
                continue
            }
            if (!(t[UL_coo.X][UL_coo.Y] == "M" && t[DR_coo.X][DR_coo.Y] == "S")){ 
                if (!(t[UL_coo.X][UL_coo.Y] == "S" && t[DR_coo.X][DR_coo.Y] == "M")){
                    continue
                }
            }
            if (!(t[UR_coo.X][UR_coo.Y] == "M" && t[DL_coo.X][DL_coo.Y] == "S")){
                if (!(t[UR_coo.X][UR_coo.Y] == "S" && t[DL_coo.X][DL_coo.Y] == "M")){
                    continue
                }
            }
            res += 1

            }
        }
    return res
}
func main(){
    fmt.Println(part2(parser()))
}
