package main

	import(
	"fmt"
	_ "embed"
	//"slices"
	"strings"
)

//go:embed input.txt
var inputDay string

func parser()[][]byte{
	res := make([][]byte,0)
	lines := strings.Split(strings.TrimRight(inputDay,"\n"),"\n")
	for _,v := range lines{
		indi := strings.Split(v,"")
		l := make([]byte,0)
		for _,by := range indi{
			l = append(l, by[0])
		}
		res = append(res, l)
	}
	return res

}

type coo_t struct{
	x int;
	y int;
}

func nb_nei(grid [][]byte,pos coo_t)int{
	res := 0
	if pos.x != 0{
		if grid[pos.x-1][pos.y] == '@' {res++}
	}
	if pos.x != len(grid[0])-1{
		if grid[pos.x+1][pos.y] == '@' {res++}
	}
	if pos.y != len(grid)-1{
		if grid[pos.x][pos.y+1] == '@' {res++}
	}
	if pos.y != 0{
		if grid[pos.x][pos.y-1] == '@' {res++}
	}
	if pos.x != 0 && pos.y != len(grid)-1{
		if grid[pos.x-1][pos.y+1] == '@' {res++}
	}
	if pos.x != len(grid[0])-1 && pos.y != len(grid)-1{
		if grid[pos.x+1][pos.y+1] == '@' {res++}
	}
	if pos.x != 0 && pos.y != 0{
		if grid[pos.x-1][pos.y-1] == '@' {res++}
	}
	if pos.x != len(grid[0])-1 && pos.y != 0{
		if grid[pos.x+1][pos.y-1] == '@' {res++}
	}
	return res
}

func part1(grid [][]byte)int{
	res := 0
	for i := range grid{
		for j := range grid[i]{
			if grid[i][j] == '@'{
				if nb_nei(grid,coo_t{i,j}) < 4 {res++}
			}
		}
	}
	return res
}
func part2(grid [][]byte)int{
	res := 0
	old := -1
	for old != 0{
		old = 0
		for i := range grid{
			for j := range grid[i]{
				if grid[i][j] == '@'{
					if nb_nei(grid,coo_t{i,j}) < 4 {res++;old++;grid[i][j] ='.'}
				}
			}
		}
	}
	return res
}

func main(){
	fmt.Println(parser())
	fmt.Println("Part1:",part1(parser()))
	fmt.Println("Part2:",part2(parser()))
}
