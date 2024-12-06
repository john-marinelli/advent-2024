package main

import (
	"advent-2024/utils"
	"fmt"
	"sort"
)

func GetLists(path string) ([]int, []int) {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	locs, err := utils.SplitNumericalLines(l)
	if err != nil {
		panic(err)
	}

	var list1 []int
	var list2 []int

	for _, loc := range locs {
		list1 = append(list1, loc[0])
		list2 = append(list2, loc[1])
	}

	return list1, list2

}

func partOne() {

	list1, list2 := GetLists("input.txt")

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var totalDist int

	for i := range list1 {
		var dist int
		dist = list1[i] - list2[i]
		if dist < 0 {
			dist = -1 * dist
		}
		totalDist += dist
	}

	fmt.Println(totalDist)
}

func partTwo() {
	l1, l2 := GetLists("input.txt")

	c := make(map[int]int)

	for _, i := range l2 {
		if _, ok := c[i]; ok {
			c[i] += 1
		} else {
			c[i] = 1
		}
	}

	simScore := 0

	for _, i := range l1 {
		if _, ok := c[i]; ok {
			simScore += (c[i] * i)
		}
	}

	fmt.Println(simScore)
}

func main() {
	partOne()
	partTwo()
}
