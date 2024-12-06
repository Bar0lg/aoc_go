package main

import (
	_ "embed"
	"fmt"
    "strings"
)

//go:embed test.txt
var inputDay string

var TEST [][]string

type coo_t struct {
	x   int
	y   int
	dir int
}
func parser(input string) ([][]string,coo_t){
    var start coo_t
	var l []string = strings.Split(input, "\n")
	l = l[:len(l)-1]
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
        TEST = append(TEST, strings.Split(v, "")) 
	}
    for i := range res{
        for j := range res[i]{
            if res[i][j] == "^"{
                start = coo_t{i,j,3}
                res[i][j] = "."
                TEST[i][j] = "."
            }
        }
    }
	return res,start
}
func apply_dir(coo coo_t) coo_t {
	switch coo.dir {
	case 0:
		return coo_t{coo.x, coo.y +1, coo.dir}
	case 1:
		return coo_t{coo.x + 1, coo.y, coo.dir}
	case 2:
		return coo_t{coo.x, coo.y - 1, coo.dir}
	case 3:
		return coo_t{coo.x - 1, coo.y, coo.dir}
	default:
		return coo_t{-1, -1, 0}

	}
}

func apply_rota(coo coo_t)coo_t{
    switch coo.dir{
       case 0:
		return coo_t{coo.x, coo.y -1, 1}
	case 1:
		return coo_t{coo.x - 1, coo.y, 2}
	case 2:
		return coo_t{coo.x, coo.y + 1, 3}
	case 3:
		return coo_t{coo.x + 1, coo.y, 0}
	default:
		return coo_t{-1, -1, 0} 
    }
}

func rec_path(tab [][]string, coo coo_t, seen map[coo_t]bool) {
    //fmt.Println(TEST)
	if coo.x < 0 || coo.x >= len(tab) || coo.y < 0 || coo.y >= len(tab[0]) {
		return
	}
    if tab[coo.x][coo.y] == "#"{
        rec_path(tab,apply_rota(coo),seen)
        return
    }
    //TEST[coo.x][coo.y] = "X"
	vue := seen[coo]
	if vue {
		return
	}
	seen[coo] = true
	if tab[coo.x][coo.y] == "." {
		rec_path(tab, apply_dir(coo), seen)
		return
}
}
func rec_pathp2(tab [][]string, coo coo_t, seen map[coo_t]bool) bool {
    //fmt.Println(TEST)
	if coo.x < 0 || coo.x >= len(tab) || coo.y < 0 || coo.y >= len(tab[0]) {
		return false
	}
    if tab[coo.x][coo.y] == "#"{
        return rec_pathp2(tab,apply_rota(coo),seen)
        
    }
    //TEST[coo.x][coo.y] = "X"
	vue := seen[coo]
	if vue {
		return true
	}
	seen[coo] = true
	return rec_pathp2(tab, apply_dir(coo), seen)

		

}
func only_cases(di map[coo_t]bool) int {
    //fmt.Println(TEST)
	var d map[coo_t]bool = make(map[coo_t]bool)
	for val := range di {
		d[coo_t{val.x, val.y, 0}] = true
        //TEST[val.x][val.y] = "0"
	}
    //fmt.Println(TEST)
	return len(d)
}

func only_cases2(di map[coo_t]bool) map[coo_t]bool {
    //fmt.Println(TEST)
	var d map[coo_t]bool = make(map[coo_t]bool)
	for val := range di {
		d[coo_t{val.x, val.y, 0}] = true
        //TEST[val.x][val.y] = "0"
	}
    //fmt.Println(TEST)
	return d
}
func part1(tab [][]string,start coo_t) int {
	var seen = make(map[coo_t]bool)
	rec_path(tab, start, seen)
	return only_cases(seen)
}

func part2(tab [][]string, start coo_t)int{
    var res int
   var seen = make(map[coo_t]bool)

   rec_path(tab,start,seen)
    all_path := only_cases2(seen)
   var seen2 = make(map[coo_t]bool)
   for k := range all_path{
       seen2 = make(map[coo_t]bool)
       if (k == coo_t{start.x,start.y,3}){
           continue
       }
       tab[k.x][k.y] = "#"
       if rec_pathp2(tab,start,seen2){
           TEST[k.x][k.y] = "0"
           res++
       }
       tab[k.x][k.y] = "."
   }
   fmt.Println(TEST)
   return res
   
}

func main(){
    grid,start_pos := parser(inputDay)
    fmt.Println(part2(grid,start_pos))
}
