package main

import (
	"advent-2024/utils"
	"fmt"
)

var Word = []string{"X", "M", "A", "S"}

type move func(int, int) (int, int)

func checkWord(
	pos int,
	x int,
	y int,
	dirFunc move,
	xLen int,
	yLen int,
	puzzle [][]string,
) int {
	if x > xLen || y > yLen || x < 0 || y < 0 || Word[pos] != puzzle[y][x] {
		return 0
	}

	if pos == (len(Word)-1) && puzzle[y][x] == Word[pos] {
		return 1
	}

	x, y = dirFunc(x, y)

	return checkWord(
		pos+1,
		x,
		y,
		dirFunc,
		xLen,
		yLen,
		puzzle,
	)
}

func up(x int, y int) (int, int) {
	return x, y - 1
}

func down(x int, y int) (int, int) {
	return x, y + 1
}

func left(x int, y int) (int, int) {
	return x - 1, y
}

func right(x int, y int) (int, int) {
	return x + 1, y
}

func upleft(x int, y int) (int, int) {
	return x - 1, y - 1
}

func upright(x int, y int) (int, int) {
	return x + 1, y - 1
}

func downleft(x int, y int) (int, int) {
	return x - 1, y + 1
}

func downright(x int, y int) (int, int) {
	return x + 1, y + 1
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	puzzle, err := utils.SplitLines(l)
	if err != nil {
		panic(err)
	}
	xLen := len(puzzle[0]) - 1
	yLen := len(puzzle) - 1

	res := 0

	// TODO Find a less wordy way of doing this
	for y := range puzzle {
		for x := range puzzle[y] {
			if puzzle[y][x] != "X" {
				continue
			}
			res += checkWord(
				0,
				x,
				y,
				up,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				down,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				left,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				right,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				upright,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				upleft,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				downright,
				xLen,
				yLen,
				puzzle,
			)
			res += checkWord(
				0,
				x,
				y,
				downleft,
				xLen,
				yLen,
				puzzle,
			)
		}
	}
	fmt.Println(res)
}

func xAllowed(l string) bool {
	return l == "M" || l == "S"
}

func checkX(x int, y int, xLen int, yLen int, puzzle [][]string) int {
	if x == xLen || y == yLen || x == 0 || y == 0 {
		return 0
	}
	if xAllowed(puzzle[y+1][x-1]) &&
		xAllowed(puzzle[y-1][x+1]) &&
		puzzle[y-1][x+1] != puzzle[y+1][x-1] &&
		xAllowed(puzzle[y+1][x+1]) &&
		xAllowed(puzzle[y-1][x-1]) &&
		puzzle[y-1][x-1] != puzzle[y+1][x+1] {
		return 1
	}

	return 0

}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	puzzle, err := utils.SplitLines(l)
	if err != nil {
		panic(err)
	}
	xLen := len(puzzle[0]) - 1
	yLen := len(puzzle) - 1
	res := 0

	for y := range puzzle {
		for x := range puzzle[y] {
			if puzzle[y][x] != "A" {
				continue
			}
			res += checkX(
				x,
				y,
				xLen,
				yLen,
				puzzle,
			)
		}
	}

	fmt.Println(res)

}

func main() {
	partOne()
	partTwo()
}
