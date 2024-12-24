package main

import (
    "fmt"
    "strconv"
    "strings"
    _ "embed"
)

//go:embed input.txt
var inputDay string

const(
    OR = 0
    AND = 1
    XOR = 2
)

type gate_t struct{
    gate_i1 string
    gate_i2 string
    gate_type int
}

type conn_t map[string]gate_t
type init_t map[string]int

func parser()(conn_t,init_t){
    splitted := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    init_str := splitted[0]
    init_split := strings.Split(init_str,"\n")
    res_init := make(init_t)
    for _,v := range init_split{
        gate_and_val := strings.Split(v,": ")
        gate := gate_and_val[0]
        n,_ := strconv.Atoi(gate_and_val[1])
        res_init[gate] = n
    }
    conn_splited := strings.Split(splitted[1],"\n")
    res_conn := make(conn_t)
    for _,v := range(conn_splited){
        details := strings.Split(v," ")
            gate := gate_t{}
            switch details[1]{
            case "AND":
                gate.gate_type = AND
            case "OR":
                gate.gate_type = OR
            case "XOR":
                gate.gate_type = XOR
            }
            gate.gate_i1 = details[0]
            gate.gate_i2 = details[2]
            res_conn[details[4]] = gate

    }
    return res_conn,res_init
}

func parcours(conn conn_t,init init_t,gate string,seen map[string]int)int{
    val,ok := seen[gate]
    if ok{
        return val
    }
    val,ok = init[gate]
    if ok{
        return val
    }
    connections  := conn[gate]
    in1 := parcours(conn,init,connections.gate_i1,seen)
    in2 := parcours(conn,init,connections.gate_i2,seen)
    res := 0
    switch connections.gate_type{
    case AND:
        res = in1 & in2
    case OR:
        res = in1 | in2
    case XOR:
        res = in1 ^ in2
    }
    seen[gate] = res
    return res
}

func Part1(conn conn_t,init init_t)int{
    seen := make(map[string]int)
    res := 0
    i := 0
    for true{
        new_z := ""
        if i < 10{
            new_z = "z0" + strconv.Itoa(i)

        }else{
            new_z = "z"+strconv.Itoa(i)
        }
        _,ok := conn[new_z]
        if !ok{
            break
        }
        res = res | ( parcours(conn,init,new_z,seen) << i)
        i++
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
