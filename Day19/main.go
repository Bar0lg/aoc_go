package main

import (
    "fmt"
    "strings"
    _ "embed"
)

//go:embed input.txt
var inputDay string

func parser()([]string,[]string){
    in := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    patterns := in[0]
    towels := in[1]
    return strings.Split(patterns,", "),strings.Split(towels,"\n")
}
func check(patt []string,str string, seen map[string]bool)bool{
    if seen[str]{
        return false
    }
    if str == ""{
        return true
    }
    seen[str] = true
    res := false
    for _,v := range patt{
        cutted := strings.TrimPrefix(str,v)
        if cutted != str{
           res = res || check(patt,cutted,seen) 
        }
    }
    return res
}

func Part1(patterns []string,towels []string)int{
    res := 0
    for _,v:= range towels{
        //fmt.Println("TEST:",v)
        if check(patterns,v,make(map[string]bool)){
            //fmt.Println("PROC:",v)
            res++
        }
    }
    return res
}
func checkv2(patt []string,str string, seen map[string]int)int{
    val,ok := seen[str]
    if ok{
        return val
    }
    if str == ""{
        return 1
    }
    res := 0
    for _,v := range patt{
        cutted := strings.TrimPrefix(str,v)
        if cutted != str{
           res +=checkv2(patt,cutted,seen) 
        }
    }
    seen[str] = res 
    return res
}
func Part2(patterns []string,towels []string)int{
    res := 0
    seen := make(map[string]int)
    for _,v:= range towels{
        //fmt.Println("TEST:",v)
        res +=checkv2(patterns,v,seen)
    }
    //fmt.Println(seen)
    return res
}

func main(){
    //fmt.Println(parser())
    fmt.Println("PART1:",Part1(parser()))
    fmt.Println("PART 2:",Part2(parser()))
}
