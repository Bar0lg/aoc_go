package main

import (
	"strings"
    "fmt"
)

func InToGrid(input string) [][]string {
	var l []string = strings.Split(input, "\n")
	l = l[:len(l)-1]
	var res [][]string = make([][]string, 0)
	for _, v := range l {
		res = append(res, strings.Split(v, ""))
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PrintGrid(g [][]string){
    for _,v := range g{
        for _,v2 := range v{
            fmt.Printf("%s",v2)
        }
        fmt.Printf("\n")
    }
}
func Copy_grid(g [][]string)[][]string{
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

