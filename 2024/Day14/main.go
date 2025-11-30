package main

import (
    "fmt"
    "strings"
    _ "embed"
)

//go:embed input.txt
var inputDay string

const (
    XMAX = 101
    YMAX = 103
    NUM_STEPS = 100
)

type robot_t struct{
    x int
    y int
    vx int
    vy int
}

func parser()[]robot_t{
    robots := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n")
    res := make([]robot_t,0)
    for _,v := range robots{
        var px,py,vx,vy int
        fmt.Sscanf(v,"p=%d,%d v=%d,%d",&px,&py,&vx,&vy)
        res = append(res, robot_t{px,py,vx,vy})
    }
    return res
}
func quadrant(r robot_t)int{
    if r.x < (XMAX-1)/2{
        if r.y < (YMAX-1)/2{
            return 0
        }else if r.y > (YMAX-1)/2{
            return 1
        }
        
    }else if r.x > (XMAX-1)/2{
        if r.y < (YMAX-1)/2{
            return 2
        }else if r.y > (YMAX-1)/2{
            return 3
        }
    }
    return -1
}

func abs(x int)int{
    if x < 0{
        return -x
    }
    return x
}

func apply_velo(rob robot_t,num_step int)robot_t{
    new_x := rob.x + rob.vx *num_step
    new_y := rob.y + rob.vy *num_step
    if new_x >= 0{
        new_x = new_x % XMAX
    }else{
        new_x = (XMAX - (abs(new_x) % XMAX)) % XMAX
    }
    if new_y >= 0{
        new_y = new_y % YMAX
    }else{
        new_y = (YMAX - (abs(new_y) % YMAX)) % YMAX
    }
    return robot_t{new_x,new_y,rob.vx,rob.vy}
}

func Part1(list_robot []robot_t)int{
    num_in_quad := [4]int{}
    res := 0
    for _,v := range list_robot{
        new_rob := apply_velo(v,NUM_STEPS)
        //fmt.Println(new_rob)
        quad := quadrant(new_rob)
        //fmt.Println(quad)
        if quad != -1{
            num_in_quad[quad]++
        }
    }
    res = num_in_quad[0]*num_in_quad[1]*num_in_quad[2]*num_in_quad[3]
    //fmt.Println(num_in_quad)
    return res
    
}

func print_grid(robs []robot_t){
    res := make([][]string,0)
    for i := range YMAX{
        res = append(res,make([]string,0))
    
        for range XMAX{
            res[i] = append(res[i], ".")
        }
    }
    for _,v := range robs{
        res[v.y][v.x] = "#"
    }
    fmt.Println(res)

}

func Part2(list_robot []robot_t){
    for i := range 100{
        new_robs := make([]robot_t,0)
    for _,v := range list_robot{
        new_robs = append(new_robs,apply_velo(v,2335 + 101*i))
        }
        fmt.Println(2335+101*i)
        print_grid(new_robs)
    }
}

func main(){
    //fmt.Println(parser())
    //create_grid(parser())
    fmt.Println("PART1:",Part1(parser()))
    Part2(parser())
}
