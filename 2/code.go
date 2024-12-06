package main

import (
	"advent-2024/utils"
	"fmt"
)

func GetReports() [][]int {
	sl, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	r, err := utils.SplitNumericalLines(sl)
	if err != nil {
		panic(err)
	}

	return r
}

func CheckSafe(report []int) bool {
	var sum int
	var absSum int
	for i := range len(report) - 1 {
		delta := report[i] - report[i+1]
		if utils.Abs(delta) > 3 || utils.Abs(delta) < 1 {
			return false
		}
		absSum += utils.Abs(delta)
		sum += delta
	}

	return absSum == utils.Abs(sum)
}

func partOne() {

	reports := GetReports()

	unsafe := 0

	for _, r := range reports {
		if !CheckSafe(r) {
			unsafe += 1
		}
	}

	fmt.Println(len(reports) - unsafe)
}

func partTwo() {
	reports := GetReports()

	safe := 0

out:
	for _, r := range reports {
		if CheckSafe(r) {
			safe += 1
			continue
		}
		for i := range len(r) {
			if CheckSafe(utils.RemoveEle(r, i)) {
				safe += 1
				continue out
			}
		}
	}

	fmt.Println(safe)
}

func main() {
	partOne()
	partTwo()
}
