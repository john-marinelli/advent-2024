package main

import (
	"advent-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// It's midnight and I'm playing catch up don't judge the time complexity

func splitRules(rules []string) [][]string {
	var res [][]string
	for _, r := range rules {
		res = append(res, strings.Split(r, "|"))
	}

	return res
}

func splitOrders(orders []string) [][]string {
	var res [][]string
	for _, r := range orders {
		res = append(res, strings.Split(r, ","))
	}

	return res
}

func makeRules(rules [][]string) map[string]map[string]bool {
	order := make(map[string]map[string]bool)

	for _, r := range rules {
		if order[r[1]] == nil {
			order[r[1]] = map[string]bool{}
		}
		order[r[1]][r[0]] = true
	}

	return order
}

func checkOrder(order []string, bfr map[string]map[string]bool) int {
	for i := range order {
		for j := i + 1; j < len(order); j++ {
			if utils.IsIn(order[j], bfr[order[i]]) {
				return 0
			}
		}
	}

	res, err := strconv.Atoi(order[len(order)/2])
	if err != nil {
		panic(err)
	}

	return res
}

func incorrect(order []string, bfr map[string]map[string]bool) bool {
	for i := range order {
		for j := i + 1; j < len(order); j++ {
			if utils.IsIn(order[j], bfr[order[i]]) {
				return true
			}
		}
	}

	return false
}

func reorder(order []string, bfr map[string]map[string]bool) int {
	// Ew gross
	for i := len(order) - 1; i >= 1; i-- {
		for j := i; j >= 0; j-- {
			if utils.IsIn(order[i], bfr[order[j]]) {
				tmp := order[i]
				order[i] = order[j]
				order[j] = tmp
			}
		}
	}

	res, err := strconv.Atoi(order[len(order)/2])
	if err != nil {
		panic(err)
	}

	return res
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	t, b := utils.SplitPage(l)
	rules := splitRules(t)
	bfr := makeRules(rules)
	ords := splitOrders(b)
	res := 0

	for _, o := range ords {
		res += checkOrder(o, bfr)
	}

	fmt.Println(res)
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	t, b := utils.SplitPage(l)
	rules := splitRules(t)
	bfr := makeRules(rules)
	ords := splitOrders(b)
	res := 0

	for _, o := range ords {
		if incorrect(o, bfr) {
			res += reorder(o, bfr)
		}
	}

	fmt.Println(res)
}

func main() {
	partOne()
	partTwo()
}
