package main

import (
	"advent-2024/utils"
	"strconv"
	"strings"
)

type Eq struct {
	A int
	O []int
}

func splitEqs(lines []string) []Eq {
	var eqs []Eq
	for _, l := range lines {
		spl := strings.Split(l, ":")
		ospl := strings.Fields(spl[1])
		var o []int
		for _, n := range ospl {
			c, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			o = append(o, c)
		}
		a, err := strconv.Atoi(spl[0])
		if err != nil {
			panic(err)
		}
		eqs = append(eqs, Eq{
			A: a,
			O: o,
		})
	}

	return eqs
}

func isEven(n int) bool {
	return (n % 2) == 0
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	eqs := splitEqs(l)
	var add []int

	for _, eq := range eqs {
	}

}

func main() {

}
