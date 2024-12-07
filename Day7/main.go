package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input.txt
var InputDay string


func to_int(t [][]string) [][]int {
	var res [][]int = make([][]int, 0)
	for i, v := range t {
		res = append(res, make([]int, 0))
		for _, s := range v {
			var n1, _ = strconv.Atoi(s)
			res[i] = append(res[i], n1)
		}
	}
	return res
}

func parser()([]int,[][]int){
    var equas []string = strings.Split(InputDay,"\n")
    equas = equas[:len(equas)-1]
    var results []int = make([]int, 0)
    var nums [][]string = make([][]string, 0)
    for _,v := range equas{
        tmp := strings.Split(v,":")
        n1,_ := strconv.Atoi(tmp[0])
        results = append(results, n1)
        nums = append(nums, strings.Split(tmp[1]," ")[1:])
    }
    return results, to_int(nums)

}

func rec_resolver(before int, numbers []int, goal int)bool{
    if (len(numbers) == 0 && before == goal){
        return true
    }
    if before > goal || len(numbers) == 0{
        return false
    }
    var res bool = false
    res = rec_resolver(before + numbers[0],numbers[1:],goal)
    res = res || rec_resolver(before * numbers[0],numbers[1:],goal)
    return res
}
func rec_resolver_part2(before int, numbers []int, goal int)bool{
    if (len(numbers) == 0 && before == goal){
        return true
    }
    if before > goal || len(numbers) == 0{
        return false
    }
    var res bool = false
    res = rec_resolver_part2(before + numbers[0],numbers[1:],goal)
    res = res || rec_resolver_part2(before * numbers[0],numbers[1:],goal)
    s1 := strconv.Itoa(before)
    s2 := strconv.Itoa(numbers[0])
    concat,_ := strconv.Atoi(s1 +s2)
    res = res || rec_resolver_part2(concat,numbers[1:],goal)
    return res
}
func part1(results []int,nums [][]int)int{
    var res int
    for i := range(results){
        if rec_resolver(nums[i][0],nums[i][1:],results[i]){
            res += results[i]
        }
    }
    return res
}
func part2(results []int,nums [][]int)int{
    var res int
    for i := range(results){
        if rec_resolver_part2(nums[i][0],nums[i][1:],results[i]){
            res += results[i]
        }
    }
    return res
}
func main(){
    fmt.Println(part2(parser()))
}
