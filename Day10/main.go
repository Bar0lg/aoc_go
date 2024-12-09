package main

import(
    "fmt"
    _ "embed"
    "strconv"
    "strings"
)

//go:embed input.txt
var inputDay string

func parser()[]int{
    var disk = strings.Split(inputDay,"")
    disk = disk[:len(disk)-1]
    var res []int = make([]int, 0)
    for i := range disk{
        n,_ := strconv.Atoi(disk[i])
        res= append(res, n)
    }
    return res
}

func sum_n(n int)int{
    return (n*(n+1))/2
}

func checksumPart1(disk []int) int{
    var res = 0
    var disk_index = 0
    var disk_index_back = len(disk) -1
    var num_back = 0
    if (disk_index_back %2 == 1){
        num_back = disk[disk_index_back-1]
    }else
    {
        num_back = disk[disk_index_back]
    }
    for i := range disk{
        if i % 2 == 0{
            for k := range(disk[i]){
                res += (disk_index+k)*(i/2)
                fmt.Println(disk_index+k,"*",i/2)
            }        
            disk_index += disk[i]
        }else{
            num_to_replace := disk[i]
            for num_to_replace != 0{
            number_taken := min(num_to_replace,num_back)
            for k := range(number_taken){
                res += (disk_index+k)*(disk_index_back/2)
                fmt.Println(disk_index+k,"*",disk_index_back/2)
            }
            //fmt.Println(disk_index)
            fmt.Println(i,num_to_replace,number_taken,disk_index_back,num_back)
            //fmt.Println(disk_index)
            if number_taken == num_back{
                disk_index_back  = max(i+1,disk_index_back-2)
                
                num_back = disk[disk_index_back]

            }else{
                num_back -= num_to_replace
            }
            num_to_replace -= number_taken
            disk_index += number_taken

        }
        if disk_index_back == i+1{
            break
        }
        //fmt.Println("RES:",res)
        
    }
}
return res
}
func main(){
    fmt.Println(parser())
    fmt.Println(checksumPart1(parser()))
}
