package main

import (
	"advent-2024/utils"
	"fmt"
)

func parse(lines []string) ([][]int, [][]int) {
	btm := 0
	keys := [][]int{}
	locks := [][]int{}
	for i, line := range lines {
		if line != "" && i != len(lines)-1 {
			continue
		}
		var cap int
		if i == len(lines)-1 {
			cap = i + 1
		} else {
			cap = i
		}

		curr := utils.ParseCartesian(lines[btm:cap])
		counts := []int{0, 0, 0, 0, 0}
		if curr[0][0] == "#" {
			for _, row := range curr[1:] {
				for j, char := range row {
					if char == "#" {
						counts[j] += 1
					}
				}
			}
			keys = append(keys, counts)
		} else {
			for i := len(curr) - 2; i >= 0; i-- {
				for j, char := range curr[i] {
					if char == "#" {
						counts[j] += 1
					}
				}
			}
			locks = append(locks, counts)
		}
		btm = i + 1
	}

	return locks, keys
}

func checkLock(lock []int, key []int, max int) bool {
	for i := range lock {
		if lock[i]+key[i] > max {
			return false
		}
	}

	return true
}

func checkCombos(locks [][]int, keys [][]int, max int) int {
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if checkLock(lock, key, max) {
				total += 1
			}
		}
	}

	return total
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	locks, keys := parse(l)
	res := checkCombos(locks, keys, 5)
	fmt.Println(res)
}

func main() {
	partOne()
}
