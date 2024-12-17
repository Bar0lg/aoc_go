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
    Reg_A uint64
    Reg_B uint64
    Reg_C uint64
    Programm []uint64
    PC uint64

}

func parser()VM_t{
    reg_and_code := strings.Split(strings.TrimSuffix(inputDay,"\n"),"\n\n")
    regs := reg_and_code[0]
    regs_split := strings.Split(regs,"\n")
    var rega uint64
    var regb uint64
    var regc uint64
    fmt.Sscanf(regs_split[0],"Register A: %d",&rega)
    fmt.Sscanf(regs_split[1],"Register B: %d",&regb)
    fmt.Sscanf(regs_split[2],"Register C: %d",&regc)
    code := reg_and_code[1]
    code_splitted := strings.Split(strings.TrimPrefix(code,"Program: "),",")
    fmt.Println(code_splitted)
    code_int := make([]uint64,0)
    for _,v := range code_splitted{
        n1,_ := strconv.Atoi(v)
        code_int = append(code_int, uint64(n1))
    }
    return VM_t{rega,regb,regc,code_int,0}
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

func adv(vm *VM_t,num uint64,denum uint64){
    vm.Reg_A = num/denum
    vm.PC += 2
}
func bxl(vm *VM_t,num1 uint64,num2 uint64){
    vm.Reg_B = num1 ^num2
    vm.PC += 2
}
func bst(vm *VM_t,num uint64){
    vm.Reg_B = num % 8
    vm.PC += 2
}
func jnz(vm *VM_t,num uint64){
    if vm.Reg_A != 0{
        vm.PC = num
        return
    }
    vm.PC += 2
}
func bxc(vm *VM_t,num1 uint64,num2 uint64){
    vm.Reg_B = num1 ^num2
    vm.PC += 2
}
func out(vm *VM_t,num uint64){
    fmt.Printf("%d,",num % 8)
    vm.PC += 2
}
func bdv(vm *VM_t,num uint64,denum uint64){
    vm.Reg_B =  num/denum
    vm.PC += 2

}
func cdv(vm *VM_t,num uint64,denum uint64){
    vm.Reg_C = num/denum
    vm.PC += 2
}
func exp(x uint64,n uint64,max_i uint64)uint64{
    if n == 0{
        return 1
    }
    if n %2 == 0{
        tmp := exp(x,n/2,max_i)
        if tmp >= max_i{
            return max_i +1
        }
        return tmp*tmp
    }
    tmp := exp(x,(n-1)/2,max_i)
    if tmp >= max_i{
        return max_i + 1
    }
    return x * tmp*tmp
}

func combo(vm *VM_t,ope uint64)uint64{
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
        adv(vm,vm.Reg_A,exp(2,combo(vm,opperand),vm.Reg_A))
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
        bdv(vm,vm.Reg_A,exp(2,combo(vm,opperand),vm.Reg_A))
    case CDV:
        cdv(vm,vm.Reg_A,exp(2,combo(vm,opperand),vm.Reg_A))
    default:
        fmt.Println("ERROR")
    }
    run(vm)
}

func Part1(vm VM_t){
    run(&vm)
    fmt.Printf("\n")
    return
}
func Part2(){
    return
}

func main(){
    fmt.Println(parser())
    Part1(parser())
    //fmt.Println("PART 2:",Part2(parser()))
}
