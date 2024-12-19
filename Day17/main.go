package main

import (
    "fmt"
    "strconv"
    "strings"
    _ "embed"
)

//go:embed input.txt
var inputDay string

type VM_t struct{
    Reg_A int
    Reg_B int
    Reg_C int
    Programm []int
    PC int
    outupt []int

}

func parser()VM_t{
    reg_and_code := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    regs := reg_and_code[0]
    regs_split := strings.Split(regs,"\n")
    var rega int
    var regb int
    var regc int
    fmt.Sscanf(regs_split[0],"Register A: %d",&rega)
    fmt.Sscanf(regs_split[1],"Register B: %d",&regb)
    fmt.Sscanf(regs_split[2],"Register C: %d",&regc)
    code := reg_and_code[1]
    code_splitted := strings.Split(strings.TrimPrefix(code,"Program: "),",")
    code_int := make([]int,0)
    for _,v := range code_splitted{
        n1,_ := strconv.Atoi(v)
        code_int = append(code_int, int(n1))
    }
    return VM_t{rega,regb,regc,code_int,0,make([]int, 0)}
}

const (
    ADV = 0
    BXL = 1
    BST = 2
    JNZ = 3
    BXC = 4
    OUT = 5
    BDV = 6
    CDV = 7
)

func adv(vm *VM_t,num int,denum int){
    vm.Reg_A = num >> denum
    vm.PC += 2
}
func bxl(vm *VM_t,num1 int,num2 int){
    vm.Reg_B = num1 ^num2
    vm.PC += 2
}
func bst(vm *VM_t,num int){
    vm.Reg_B = num % 8
    vm.PC += 2
}
func jnz(vm *VM_t,num int){
    if vm.Reg_A != 0{
        vm.PC = num
        return
    }
    vm.PC += 2
}
func bxc(vm *VM_t,num1 int,num2 int){
    vm.Reg_B = num1 ^num2
    vm.PC += 2
}
func out(vm *VM_t,num int){
    //fmt.Printf("%d,",num%8)
    vm.outupt = append(vm.outupt, num%8)
    vm.PC += 2
}
func bdv(vm *VM_t,num int,denum int){
    vm.Reg_B =  num >> denum
    vm.PC += 2

}
func cdv(vm *VM_t,num int,denum int){
    vm.Reg_C = num >> denum
    vm.PC += 2
}

func combo(vm *VM_t,ope int)int{
    switch ope{
    case 0:
        return 0
    case 1:
        return 1
    case 2:
        return 2
    case 3:
        return 3
    case 4:
        return vm.Reg_A
    case 5:
        return vm.Reg_B
    case 6:
        return vm.Reg_C
    default:
        fmt.Println("ERROR")
        return 999

    }
}

func run(vm *VM_t){
    if int(vm.PC) >= len(vm.Programm)-1{
        return
    }
    code := vm.Programm[vm.PC]
    opperand := vm.Programm[vm.PC+1]
    switch code{
    case ADV:
        adv(vm,vm.Reg_A,combo(vm,opperand))
    case BXL:
        bxl(vm,vm.Reg_B,opperand)
    case BST:
        bst(vm,combo(vm,opperand))
    case JNZ:
        jnz(vm,opperand)
    case BXC:
        bxc(vm,vm.Reg_B,vm.Reg_C)
    case OUT:
        out(vm,combo(vm,opperand))
    case BDV:
        bdv(vm,vm.Reg_A,combo(vm,opperand))
    case CDV:
        cdv(vm,vm.Reg_A,combo(vm,opperand))
    default:
        fmt.Println("ERROR")
    }
    run(vm)
}

func Part1(vm VM_t){
    run(&vm)
    fmt.Println("Part1: ",vm.outupt)
    return
}

func compare_strings(i1 []int,i2 []int)bool{
    if len(i1) != len(i2){
        return false
    }
    for i := range(i1){
        if i1[i] != i2[i]{
            return false
        }
    }
    return true
}
func Part2(vm VM_t)int{
    var res int = 0
    var last []int
    for range vm.Programm{
        last = append(last,-1)
    }
    found := false
    num_index := 0
    i := 0
    for true {
        found = false
        for i < 8{
            vm.outupt = make([]int, 0)
            vm.Reg_A = res | int(i)
            vm.Reg_B = 0
            vm.Reg_C = 0
            vm.PC = 0
            run(&vm)
            if compare_strings(vm.outupt,vm.Programm[len(vm.Programm)-num_index-1:]){
                num_index++
                res = res | i
                if num_index == len(vm.Programm){
                    return res
                }
                res = res << 3
                i = 0
                found = true
                break
            }
            i++
        }
        if !found{
            res = res>> 3
            i = (res & 0b111) +1
            res = (res & (^0b111))
            num_index--
        }
    
    }
    return res
    }

func main(){
    //fmt.Println(parser())
    Part1(parser())
    fmt.Println("PART 2:",Part2(parser()))
}
