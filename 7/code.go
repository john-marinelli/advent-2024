package main

import (
	"advent-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Op func(a int, b int) int

var OpMap map[rune]Op = map[rune]Op{
	'0': func(a int, b int) int { return a + b },
	'1': func(a int, b int) int { return a * b },
	'2': func(a int, b int) int {
		return utils.StrToInt(strconv.Itoa(a) + strconv.Itoa(b))
	},
}

func getCalcs(lines []string, base int) ([]int, [][]int, []int) {
	ans := []int{}
	eq := [][]int{}
	combs := []int{}
	for i, line := range lines {
		tmp := strings.Split(line, ":")
		ans = append(ans, utils.StrToInt(tmp[0]))
		e := strings.Fields(tmp[1])
		ops := int(math.Pow(float64(base), float64(len(e)-1)))
		combs = append(combs, ops)
		eq = append(eq, []int{})
		for _, c := range e {
			eq[i] = append(eq[i], utils.StrToInt(c))
		}
	}

	return ans, eq, combs
}

func checkValid(a int, eq []int, comb int, base int) int {
	for j := range comb {
		curr := strconv.FormatInt(int64(j), base)
		binStr := fmt.Sprintf(`%0*s`, int(math.Log(float64(comb))/math.Log(float64(base))), curr)
		k := 1
		res := eq[0]
		for _, op := range binStr {
			res = OpMap[op](res, eq[k])
			k += 1
		}
		if a == res {
			return a
		}
	}

	return 0
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	ans, eqs, combs := getCalcs(l, 2)
	total := 0
	for i := range ans {
		total += checkValid(ans[i], eqs[i], combs[i], 2)
	}
	fmt.Println("Part one: ", total)
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	ans, eqs, combs := getCalcs(l, 3)
	total := 0
	for i := range ans {
		total += checkValid(ans[i], eqs[i], combs[i], 3)
	}
	fmt.Println("Part two: ", total)
}

func main() {
	partOne()
	partTwo()
}
