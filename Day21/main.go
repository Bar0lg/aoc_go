package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

type seq []rune
type parcours func(rune,rune)rune

func parser()[]seq{
    codes := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    res := make([]seq,0)
    for _,code := range codes{
        tmp := make(seq,0)
        for _,c := range code{
            tmp = append(tmp, rune(c))
        }
        res = append(res, tmp)
    }
    return res
}

const ERROR = 'Z'
var ALL_DIRS []rune = []rune{'<','^','>','v'}

func is_in(l []rune,e rune)bool{
    for _,v := range l{
        if v == e{
            return true
        }
    }
    return false
}
func cpy(l []rune)[]rune{
    res := make([]rune,len(l))
    copy(res,l)
    return res
}

func min_length(l []seq)int{
    if len(l) == 0{
        return 0
    }
    res := len(l[0])
    for _,v := range l{
        res = min(res,len(v))
    }
    return res
}

func print_seq(s seq){
    for _,v := range s{
        fmt.Printf("%s",string(v))
    }
    fmt.Printf("\n")
}
func seq_to_str(s seq)string{
    res := ""
    for _,v := range s{
        res = res +string(v)
    }
    return res
}
func print_all_seqs(s []seq){
    fmt.Printf("--------\n")
    for _,v := range s{
        print_seq(v)
    }
    fmt.Printf("--------\n")
}

var path_dico map[string][]seq = make(map[string][]seq)

func gen_path(start rune,end rune,seen []rune,fn parcours)[]seq{
    if is_in(seen,start){
        return []seq{nil}
    }
    if start == end{
        return []seq{{'A'}}
    }
    res := make([]seq,0)
    seen = append(seen,start)
    for _,dir := range ALL_DIRS{
        next_button := fn(start,dir)
        if next_button == ERROR{
            continue
        }
        next_steps := gen_path(fn(start,dir),end,cpy(seen),fn)
        min_l := min_length(next_steps)
        for _,next := range next_steps{
            if (next == nil){
                continue
            }
            if len(next) != min_l{
                continue
            }
            tmp := make(seq,0)
            tmp = append(tmp,dir)
            tmp = append(tmp, next...)
            res = append(res, tmp)
        }

    }
    return trim_largest(res)

}

func pad_to_pad(to_do seq,start rune,pad_fn parcours,seen map[string][]seq)[]seq{
    if len(to_do) == 0{
        return nil
    }
    str := seq_to_str(to_do)
    if seen[str] != nil{
        return seen[str]
    }
    res := gen_path(start,to_do[0],make([]rune, 0),pad_fn)
    //print_all_seqs(res)
    new_res := make([]seq,0)
    for i := range res{
        next := pad_to_pad(to_do[1:],to_do[0],pad_fn,seen)
        if next == nil{
            new_res = append(new_res, res[i])
        }
        for j := range next{
            tmp := make(seq,0)
            tmp= append(tmp, res[i]...)
            tmp = append(tmp, next[j]...)
            new_res = append(new_res, tmp)
        }
    }
    seen[str] = trim_largest(new_res)
    return seen[str]
}
func pad_to_padv2(to_do seq,start rune,seen map[string]seq)seq{
    if len(to_do) == 0{
        return nil
    }
    str := seq_to_str(to_do)
    if seen[str] != nil{
        return seen[str]
    }
    //print_all_seqs(res)
    res := dir_to_dir(start,to_do[0])
    next := pad_to_padv2(to_do[1:],to_do[0],seen)
    res = append(res, next...)
    seen[str] = res
    return seen[str]
}
func trim_largest(s []seq)[]seq{
    if len(s) == 0{
        return nil
    }
    min_i := min_length(s)
    res := make([]seq,0)
    for _,v := range s{
        if len(v) == min_i{
            res = append(res, v)
        }
    }
    return res

}
func num_code(s seq)int{
    var str string = ""
    for _,v := range s{
        str = str + string(v)
    } 
    n1,_ := strconv.Atoi(str[:len(str)-1])
    return n1
}
func Part1(codes []seq)int{
    var res int = 0
    for _,code  := range codes{
        var all_seqs []seq = codes
        var new_all_seqs []seq = make([]seq, 0)

        new_all_seqs = pad_to_pad(code,'A',Apply_dir_numpad,make(map[string][]seq))

        all_seqs = new_all_seqs
        new_all_seqs = make([]seq, 0)

        for _,v := range all_seqs{
            new_all_seqs = append(new_all_seqs, pad_to_pad(v,'A',Apply_dir_dir,make(map[string][]seq))...)
        }
        all_seqs = new_all_seqs
        new_all_seqs = make([]seq, 0)

        for _,v := range all_seqs{
            new_all_seqs = append(new_all_seqs, pad_to_pad(v,'A',Apply_dir_dir,make(map[string][]seq))...)
        }
        min_len := len(new_all_seqs[0])
        for _,v := range new_all_seqs{
            min_len = min(min_len,len(v))
        }
        num := num_code(code)
        res += min_len*num
    }
    return res
    
}




//func Part2(){
//    return
//}

func main(){
    //fmt.Println(num_code(parser()[0]))
    fmt.Println("PART1:",Part1(parser()))
    //fmt.Println("PART 2:",Part2(parser()))
}
