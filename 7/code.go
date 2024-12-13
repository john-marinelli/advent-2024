package main

import (
	"advent-2024/utils"
	"strconv"
	"strings"
)

type Op func(a int, b int) int

var OpMap map[string]Op = map[string]Op{
	"0": func(a int, b int) int { return a + b },
	"1": func(a int, b int) int { return a * b },
}

func getCalcs(lines []string, base int) ([]int, [][]int, []string) {
	ans := []int{}
	eq := [][]int{}
	combs := []string{}
	for i, line := range lines {
		tmp := strings.Split(line, ":")
		ans = append(ans, utils.StrToInt(tmp[0]))
		e := strings.Fields(tmp[1])
		ops := int64(len(e)) - 1
		combs = append(combs, strconv.FormatInt(ops, 2))
		eq = append(eq, []int{})
		for _, c := range e {
			eq[i] = append(eq[i], utils.StrToInt(c))
		}
	}

	return ans, eq, combs
}

func checkCalcs()

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	ans, eqs, combs := getCalcs(l, 2)
}

func main() {

}
