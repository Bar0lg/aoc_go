package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
)

//go:embed input.txt
var inputDay string

func parser()([]byte,[]int){
	res1 := make([]byte,0);
	res2 := make([]int,0);
	lines := strings.Split(strings.TrimRight(inputDay,"\n"),"\n");
	for _,v := range lines{
		res1  = append(res1, v[0])
		in,_ := strconv.Atoi(strings.TrimLeft(v,"RL"))
		res2 = append(res2, in)
	}
	return res1,res2

}

func part1(sens []byte,rota []int)int{
	tick := 50
	sens_rata := 1
	res := 0
	for i := range sens{
		if sens[i] == 'L'{
			sens_rata = -1
		}else{
			sens_rata = 1
		}
		tick += sens_rata* rota[i]
		for tick < 0 || tick > 99{
			if tick > 99{
				tick += -100
			}
			if tick < 0{
				tick += 100
			}
		}
		if (tick == 0){
			res++
		}

	}
	return res
}

func abs(x int)int{
	if x> 0{
		return x
	}
	return -x
}

func part2(sens []byte,rota []int)int{
	tick := 50
	sens_rata := 1
	res := 0
	for i := range sens{
		if sens[i] == 'L'{
			sens_rata = -1
			if tick == 0{
				res--
			}
		}else{
			sens_rata = 1
		}
		tick += sens_rata* rota[i]
		for tick < 0 || tick > 99{
			if tick != 100{
				res++
			}
			if tick > 99{
				tick += -100
			}
			if tick < 0{
				tick += 100
			}
		}
		if (tick == 0){
			res++
		}

	}
	return res
}

func main(){
	fmt.Println("Part1:",part1(parser()))
	fmt.Println("Part2:",part2(parser()))
}
