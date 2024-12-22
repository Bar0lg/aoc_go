package main

import (
    "fmt"
    "strconv"
    "strings"
    _ "embed"
)

//go:embed test.txt
var inputDay string

func parser()[]int{
    res := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    res_int := make([]int,0)
    for _,v := range res{
        n1,_ := strconv.Atoi(v)
        res_int = append(res_int, n1)
    }
    return res_int
}

func passage(n uint64)uint64{
    sec_num := n
    sec_64 := sec_num << 6
    sec_num = sec_64 ^ sec_num
    sec_num = sec_num % 16777216
    sec_32 := sec_num >> 5
    sec_num = sec_32 ^ sec_num
    sec_num = sec_num % 16777216
    sec_2048 := sec_num << 11
    sec_num = sec_2048 ^ sec_num
    sec_num = sec_num % 16777216
    return sec_num
}

func Part1(init []int)uint64{
    var res uint64 = 0
    for _,v := range init{
        v_64 := uint64(v)
        for range 2000{
            v_64 = passage(v_64)
        }
        res += v_64
        fmt.Println(v,v_64)
    }
    return res
}

func Part2(init []int)int{
    var res int = 0
    var true_res int= 0
    byers := make([][]int,0)
    for i,v := range init{
        byers = append(byers, make([]int, 0))
        byers[i] = append(byers[i], v)
        for k :=  range 2000{
                byers[i] = append(byers[i], int(passage(uint64(byers[i][k]))))
        }
    }

    for i:=-9;i<10;i++{
        for j:=-9;j<10;j++{
            for k:=-9;j<10;k++{
                for l:=-9;l<10;l++{
                    for _,by := range(byers){
                        changes := [4]int{}
                        changes[0] = by[1] - by[0]
                        changes[1] = by[2] - by[1]
                        changes[2] = by[3] - by[2]
                        changes[3] = by[4] - by[3]
                        for p:=4;p<len(by)-1;p++{
                            if i == changes[0] && j == changes[1] && k == changes[2] && l == changes[3]{
                                fmt.Println("PROC:",i)
                                res += by[p] % 10
                                break
                            }
                            changes[0] = changes[1]
                            changes[1] = changes[2]
                            changes[2] = changes[3]
                            changes[3] = by[p+1] - by[p]
                        }
                    }
                    true_res = max(true_res,res)
                    res = 0

                }
            }
        }
    }
    return true_res

}

func main(){
    //fmt.Println(parser())
    //fmt.Println("PART1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
