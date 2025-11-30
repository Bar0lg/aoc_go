package main

import (
    "fmt"
    "strings"
    _ "embed"
    "sort"
)

//go:embed test.txt
var inputDay string

type network_t map[string]([]string)

func parser()network_t{
    lines := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    var res network_t = make(network_t)
    for _,link := range lines{
        link_splitted := strings.Split(link,"-")
        res[link_splitted[0]] = append(res[link_splitted[0]], link_splitted[1])
        res[link_splitted[1]] = append(res[link_splitted[1]], link_splitted[0])
    }
    return res
}

func in_commons(l1 []string,l2 []string)[]string{
    var res []string = make([]string, 0)
    for _,v1 := range l1{
        for _,v2 := range l2{
            if v1 == v2{
                res = append(res, v1)
            }
        }
    }
    return res
}

func sort_3(s [3]string)[3]string{
    slice := make([]string,3)
    slice[0] = s[0]
    slice[1] = s[1]
    slice[2] = s[2]
    sort.Strings(slice)
    res := [3]string{}
    res[0] = slice[0]
    res[1] = slice[1]
    res[2] = slice[2]
    return res
}

func Part1(net network_t)int{
    var res int = 0
    var seen map[[3]string]bool = make(map[[3]string]bool)
    for node,neis := range net{
        for _,nei := range neis{
            commons := in_commons(net[nei],neis)
            for _,third := range commons{
                if (seen[sort_3([3]string{node,nei,third})]){
                    continue
                }
                if node[0] == 't' || nei[0] == 't' || third[0] == 't'{
                    res += 1
                    seen[sort_3([3]string{node,nei,third})] = true
                }
            }

            }
            
        }
    return res
}

func parcours(net network_t,node string,commons []string,seen map[string]bool)(int,[]string){
    if seen[node]{
        return 0,nil
    }
    seen[node] = true
    new_common := in_commons(commons,net[node])
    if len(new_common) == 0{
        return 1,[]string{node}
    }
    max_res := 0
    max_prevs := make([]string,0)
    for _,nei := range(new_common){
        new_max,new_prev := parcours(net,nei,new_common,seen)
        if new_max > max_res{
            max_res = new_max
            max_prevs = new_prev

        }
    }
    res := []string{node}
    res = append(res, max_prevs...)
    return 1+max_res ,res
}

func Part2(net network_t){
    max_network := make([]string,0)
    connected_num_max := 0
    for comp,neis := range net{
        connected_num,connected_net := parcours(net,comp,neis,make(map[string]bool))
        if connected_num > connected_num_max{
            connected_num_max = connected_num
            max_network = connected_net
        }

    }
    sort.Strings(max_network)
    for _,v := range max_network{
        fmt.Printf("%s,",v)
    }
    fmt.Printf("\n")
}

func main(){
    fmt.Println(parser())
    fmt.Println("PART1:",Part1(parser()))
    fmt.Printf("Part2:")
    Part2(parser())
}
