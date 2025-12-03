package main

	import(
	"fmt"
	_ "embed"
	"strings"
	"strconv"
)

//go:embed input.txt
var inputDay string

func parser()[][]int{
	res := make([][]int,0)
	ranges := strings.Split(strings.TrimRight(inputDay,"\n"),",")
	for _,v := range ranges{
		r := strings.Split(v,"-")
		i1,_ := strconv.Atoi(r[0])
		i2,_ := strconv.Atoi(r[1])
		res_1 := make([]int,0)
		res_1 = append(res_1, i1)
		res_1 = append(res_1, i2)
		res = append(res, res_1)
	}
	return res
	
}

func invalid(x int)bool{
	sx := strconv.Itoa(x)
	n := len(sx)
	if (n %2 == 1){
		return false;
	}
	if (sx[:n/2] == sx[n/2:]){
		return true
	}
	return false
}
func invalid2(x int)bool{
	sx := strconv.Itoa(x)
	n := len(sx)
	test := true
	for i:=2;i<n+1;i++{
		if n%i != 0{
			continue
		}
		test = true
		str_tst := sx[:n/i]
		for j:=1;j<i+1;j++{
			test = test && (str_tst == sx[(n*(j-1))/i:(n*j)/i])
		}
		if test{
			return true
		}
	}

	return false
}

func part1(l [][]int)int{
	res := 0
	for _,r := range l{
		for i := r[0];i<r[1]+1;i++{
			if (invalid(i)){
				res += i
			}
		}
	}
	return res
}

func part2(l [][]int)int{
	res := 0
	for _,r := range l{
		for i := r[0];i<r[1]+1;i++{
			if (invalid2(i)){
				res += i
			}
		}
	}
	return res
}

func main(){
	fmt.Println("Part1:",part1(parser()))
	fmt.Println("Part2:",part2(parser()))
}
