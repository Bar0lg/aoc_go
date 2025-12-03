package main

	import(
	"fmt"
	_ "embed"
	"strings"
	"strconv"
	"slices"
)

//go:embed input.txt
var inputDay string

func parser()[][]int{
	res := make([][]int,0)
	batts := strings.Split(strings.TrimRight(inputDay,"\n"),"\n")
	for _,b := range batts{
		batt := strings.Split(b,"")
		ba := make([]int,0)
		for _,bi := range batt{
			ati,_ := strconv.Atoi(bi)
			ba = append(ba, ati)
		}
		res = append(res, ba)
	}
	return res
}

func argmax( t[] int,exclude []int)int{
	res := -1
	max_v := -1
	for i,v := range t{
		if slices.Contains(exclude,i){
			continue
		}
		if max_v < v{
			max_v = v
			res = i
		}
	}
	return res

}

func big(ba []int)int{
	res := 0;
	m1 := argmax(ba,[]int{-1})
	m2 := 0
	if m1 != len(ba)-1{
		m2 = argmax(ba[m1:],[]int{0}) + m1
	}else{
		m2 = argmax(ba,[]int{m1})
	}
	if m1 > m2{
		res = ba[m2]*10+ba[m1]
	}else{
		res = ba[m1]*10+ba[m2]
	}
	return res


}

func Pow(i int,n int)int{
	res := 1
	for range n{
		res *= i
	}
	return res
}

func big2(ba []int)int{
	res := 0
	arrs := make([]int,0)
	nb_made := 0
	borne := 0
	for range 12{
		//fmt.Println("HIIII",arrs)
		m := 0
			//fmt.Println(ba[arrs[nb_made-1]:])

		borne = -1
		arr_cpy := make([]int,len(arrs))
		copy(arr_cpy,arrs)
		//Check borne
			{
			free := false
			for i:=len(ba)-1;i>-1;i--{
				if slices.Contains(arr_cpy,i){
					if free{
						borne = i
						break;
					}
				}else{
					free = true
				}
			}
			if borne == -1{
				borne = 0
			}
		}
		//fmt.Println("BORNE:",borne)
		//---------
		for i := range arr_cpy{
			arr_cpy[i] -= borne
		}
		m = argmax(ba[borne:],arr_cpy) + borne
		arrs = append(arrs, m)
		nb_made++
	}
	slices.Sort(arrs)
	for i,v := range arrs{
		res += ba[v] * Pow(10,11-i)
	}

	return res
}

func part1(batts [][]int)int{
	res := 0;
	for _,v := range batts{
		res += big(v);
	}
	return res


}
func part2(batts [][]int)int{
	res := 0;
	for _,v := range batts{
		res += big2(v);
	}
	return res


}

func main(){
	fmt.Println("Part1:",part1(parser()))
	fmt.Println("Part2:",part2(parser()))
}
