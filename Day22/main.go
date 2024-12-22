package main

import (
    "fmt"
    "strconv"
    "strings"
    _ "embed"
)

//go:embed input.txt
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
    }
    return res
}

func add_elem(l [4]int,e int)[4]int{
    res := [4]int{}
    res[0] = l[1]
    res[1] = l[2]
    res[2] = l[3]
    res[3] = e
    return res
}

func rand_to_banana_sold(rand []int){
    for i := range rand{
        rand[i] = rand[i] % 10
    }
}

func all_byers_conv(byers [][]int){
    for i := range byers{
        rand_to_banana_sold(byers[i])
    }
}

func Part2(init []int)int{
    var res int = 0
    byers := make([][]int,0)
    for i,v := range init{
        byers = append(byers, make([]int, 0))
        byers[i] = append(byers[i], v)
        for k :=  range 1999{
                byers[i] = append(byers[i], int(passage(uint64(byers[i][k]))))
        }
    }

    all_byers_conv(byers)
    //fmt.Println(byers)
    all_byers_deals := make([]map[[4]int]int,0)
    for _,prices := range byers{
        deals := make(map[[4]int]int)
        changes := [4]int{}
        changes = add_elem(changes,prices[1] - prices[0])
        changes = add_elem(changes,prices[2] - prices[1])
        changes = add_elem(changes,prices[3] - prices[2])
        changes = add_elem(changes,prices[4] - prices[3])
        deals[changes] = prices[4]
        for p := 5;p<len(prices);p++{
            changes = add_elem(changes,prices[p] - prices[p-1])
            _,ok := deals[changes]
            if !ok{
                deals[changes] = prices[p]
            }
        }
        all_byers_deals = append(all_byers_deals, deals)
        
    }
    all_deal_concat := make(map[[4]int]int)
    for _,byer := range all_byers_deals{
        for key,val := range byer{
            all_deal_concat[key] += val
        }
    }
    for _,val := range all_deal_concat{
        res = max(res,val)

    }
    return res
}

func main(){
    //fmt.Println(parser())
    fmt.Println("PART1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
