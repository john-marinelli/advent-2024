package main

import (
	"advent-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction func(op int, ptr int, combo []*int) int

var A *int
var B *int
var C *int
var Zero *int
var One *int
var Two *int
var Three *int
var Output []int

func initialize(a int, b int, c int) ([]*int, []Instruction) {
	A = new(int)
	B = new(int)
	C = new(int)
	*A = a
	*B = b
	*C = c
	Zero = new(int)
	One = new(int)
	Two = new(int)
	Three = new(int)
	*Zero = 0
	*One = 1
	*Two = 2
	*Three = 3
	combo := []*int{
		Zero,
		One,
		Two,
		Three,
		A,
		B,
		C,
	}
	insts := []Instruction{
		adv,
		bxl,
		bst,
		jnz,
		bxc,
		out,
		bdv,
		cdv,
	}

	return combo, insts
}

func adv(op int, ptr int, combo []*int) int {
	*A = *A / int(math.Exp2(float64(*combo[op])))
	return ptr + 2
}

func bxl(op int, ptr int, _ []*int) int {
	*B = *B ^ op
	return ptr + 2
}

func bst(op int, ptr int, combo []*int) int {
	*B = *combo[op] % 8
	return ptr + 2
}

func jnz(op int, ptr int, _ []*int) int {
	if *A == 0 {
		return ptr + 2
	}
	return op
}

func bxc(_ int, ptr int, _ []*int) int {
	*B = *B ^ *C
	return ptr + 2
}

func out(op int, ptr int, combo []*int) int {
	Output = append(Output, *combo[op]%8)
	return ptr + 2
}

func bdv(op int, ptr int, combo []*int) int {
	*B = *A / int(math.Exp2(float64(*combo[op])))
	return ptr + 2
}

func cdv(op int, ptr int, combo []*int) int {
	*C = *A / int(math.Exp2(float64(*combo[op])))
	return ptr + 2
}

func parseInput(lines []string) (int, int, int, []int) {
	nums := []int{}
	a := utils.StrToInt(strings.Trim(strings.Split(lines[0], ":")[1], " "))
	b := utils.StrToInt(strings.Trim(strings.Split(lines[1], ":")[1], " "))
	c := utils.StrToInt(strings.Trim(strings.Split(lines[2], ":")[1], " "))
	s := strings.Split(strings.TrimSpace(strings.Split(lines[4], ":")[1]), ",")
	for i := range s {
		nums = append(nums, utils.StrToInt(s[i]))
	}

	return a, b, c, nums
}

func compute(nums []int, insts []Instruction, combo []*int) {
	ptr := 0
	for ptr+1 < len(nums) {
		ptr = insts[nums[ptr]](nums[ptr+1], ptr, combo)
	}
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	a, b, c, nums := parseInput(l)
	combo, insts := initialize(a, b, c)
	compute(nums, insts, combo)
	strOutput := []string{}
	for _, o := range Output {
		strOutput = append(strOutput, strconv.Itoa(o))
	}
	fmt.Println(strings.Join(strOutput, ","))
}

func main() {
	partOne()
}
