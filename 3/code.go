package main

import (
	"advent-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func SumResult(sNums [][]string) int {

	var nums [][]int

	for _, snr := range sNums {
		var nr []int
		for _, n := range snr {
			conv, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nr = append(nr, conv)
		}
		nums = append(nums, nr)
	}

	res := 0

	for _, n := range nums {
		res += (n[0] * n[1])
	}

	return res
}

func partOne() {
	f := utils.GetFileString("input.txt")
	re := `mul\(([0-9]+),([0-9]+)\)`
	sNums := utils.MatchReExact(f, re, 2)

	fmt.Println(SumResult(sNums))

}

func partTwo() {
	f := utils.GetFileString("input.txt")
	re := `do\(\)|don't\(\)|mul\(([0-9]+),([0-9]+)\)`
	matches := utils.MatchRe(f, re)

	do := true
	res := 0

	for _, match := range matches {
		if strings.Contains(match[0], "don't") {
			do = false
		} else if strings.Contains(match[0], "do") {
			do = true
		} else if strings.Contains(match[0], "mul") && do {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			res += a * b
		}
	}

	fmt.Println(res)

}

func main() {
	partOne()
	partTwo()
}
