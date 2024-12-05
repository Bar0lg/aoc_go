package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser() ([]int, []int){
   var l1 []int = make([]int, 0)
   var l2 []int = make([]int, 0)
   var elms = strings.Split(inputDay,"\n")
   elms = elms[:len(elms)-1]
   for _,val := range(elms){
        var ints = strings.Split(val,"   ")
        var i1 , _ = strconv.Atoi(ints[0])
        var i2 , _ = strconv.Atoi(ints[1])
        l1 = append(l1,i1)
        l2 = append(l2,i2)
   }
   return l1, l2
}

func abs(x int)int{
    if (x < 0){
        return -x
    }
    return x
}

func part1(l1 []int,l2 []int) int {
    var l1_s []int= make([]int, len(l1))
    var l2_s []int= make([]int, len(l2))
    copy(l1_s,l1)
    copy(l2_s,l2)
    sort.Ints(l1_s)
    sort.Ints(l2_s)
    var res int = 0
    for i:= range(l1_s){
        res = res + abs(l1_s[i] - l2_s[i])
    }
    return res

}

func part2(l1 []int,l2 []int)int{
    var res int = 0
    for _,v1 := range(l1){
        for _,v2 := range(l2){
            if (v1 == v2){
                res += v1
            }

        }
    }
    return res
}

func main(){
    fmt.Println(part2(parser()))
}
