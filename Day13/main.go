package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var inputDay string

type coo_t struct {
	x int
	y int
}

const (
	CLAW_A = 0
	CLAW_B = 1
	PRIZE  = 2
)

func parser() [][3]coo_t {
	games := strings.Split(strings.TrimSuffix(inputDay, "\n"), "\n\n")
	res := make([][3]coo_t, 0)
	for _, v := range games {
		claws := strings.Split(v, "\n")
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(claws[CLAW_A], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(claws[CLAW_B], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(claws[PRIZE], "Prize: X=%d, Y=%d", &px, &py)
		res = append(res, [3]coo_t{{ax, ay}, {bx, by}, {px, py}})
	}
	return res
}

func apply_game(game [3]coo_t, numA int, numB int) bool {
	var co coo_t
	co.x = numA*game[CLAW_A].x + numB*game[CLAW_B].x
	co.y = numA*game[CLAW_A].y + numB*game[CLAW_B].y
	return co == game[PRIZE]

}

func Part1(games [][3]coo_t) int {
	res := 0
	for _, g := range games {
		done := false
		for i := range 101 {
			for j := range 101 {
				if apply_game(g, i, j) {
					res += 3*i + j
					//fmt.Println("RES:",i,j)
					done = true
					break
				}

			}
			if done {
				break
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part2(games [][3]coo_t) int {
	res := 0
	for _, g := range games {
		var A = g[CLAW_A]
		var B = g[CLAW_B]
		var P = coo_t{g[PRIZE].x + 10000000000000, g[PRIZE].y + 10000000000000}

		var numx int = P.y*B.x - P.x*B.y
		var denomx int = A.y*B.x - A.x*B.y

		var numy int = P.y*A.x - P.x*A.y
		var denomy int = B.y*A.x - B.x*A.y

		if abs(numx) % abs(denomx) == 0 && abs(numy) % abs(denomy) == 0 {
			res += 3*(numx/denomx) + (numy / denomy)
		}
	}
	return res
}

func main() {
	//fmt.Println(parser())
	fmt.Println("PART 1:", Part1(parser()))
	fmt.Println("PART 2:", Part2(parser()))
}
