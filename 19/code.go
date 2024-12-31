package main

import (
	"advent-2024/utils"
	"fmt"
	"strings"
)

type AdjMatrix map[string]map[string]bool

func parseInput(lines []string) ([]string, []string) {
	towels := []string{}
	tmp := strings.Split(lines[0], ",")
	for _, t := range tmp {
		towels = append(towels, strings.TrimSpace(t))
	}
	combs := []string{}
	combs = append(combs, lines[2:]...)
	return towels, combs
}

func constructGraph(towels []string) (AdjMatrix, map[string]bool) {
	return AdjMatrix{}, make(map[string]bool)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return factorial(n-1) * n
}

func getCombinations(towels []string) map[string]bool {
	perms := utils.GetPermutations(towels)
	combs := make(map[string]bool)
	for _, p := range perms {
		for i := 1; i <= len(p); i++ {
			c := strings.Join(p[:i], "")
			combs[c] = true
		}
	}

	return combs
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	towels, combs := parseInput(l)
	allCombs := getCombinations(towels)
	total := 0

	for _, c := range combs {
		if _, ok := allCombs[c]; ok {
			total += 1
		}
	}
	fmt.Println(total)
}

func main() {
	partOne()
}
