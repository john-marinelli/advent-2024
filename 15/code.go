package main

import (
	"advent-2024/utils"
	"fmt"
)

type Vec struct {
	X int
	Y int
}

type Robot struct {
	Pos Vec
}

type DBox struct {
	Char string
	Old  Vec
	New  Vec
	Prev *DBox
}

type DirFunc func(x int, y int) (int, int)

var DF map[rune]DirFunc = map[rune]DirFunc{
	'^': func(x int, y int) (int, int) {
		return x, y + 1
	},
	'v': func(x int, y int) (int, int) {
		return x, y - 1
	},
	'>': func(x int, y int) (int, int) {
		return x + 1, y
	},
	'<': func(x int, y int) (int, int) {
		return x - 1, y
	},
}

func splitDirs(lines []string) ([][]string, string) {
	i := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		i += 1
	}
	grid := lines[:i]
	d := lines[i+1:]
	dirs := ""

	wh := utils.ParseCartesian(grid)
	for _, dline := range d {
		dirs += dline
	}

	return wh, dirs
}

func findRobot(wh [][]string) *Robot {
	rx := 0
	ry := 0
	for i, line := range wh {
		for j, char := range line {
			if char == "@" {
				rx = j
				ry = i
			}
		}
	}

	return &Robot{
		Pos: Vec{
			X: rx,
			Y: ry,
		},
	}
}

func move(d rune, r *Robot, wh [][]string) {
	x, y := DF[d](r.Pos.X, r.Pos.Y)
	if wh[y][x] == "#" {
		return
	}
	if wh[y][x] == "." {
		wh[r.Pos.Y][r.Pos.X] = "."
		wh[y][x] = "@"
		r.Pos.Y = y
		r.Pos.X = x
		return
	}
	x1 := x
	y1 := y
	for wh[y1][x1] != "." && wh[y1][x1] != "#" {
		x1, y1 = DF[d](x1, y1)
	}
	if wh[y1][x1] == "#" {
		return
	}
	wh[y1][x1] = "O"
	wh[y][x] = "@"
	wh[r.Pos.Y][r.Pos.X] = "."
	r.Pos.X = x
	r.Pos.Y = y
}

func moveDouble(d rune, r *Robot, wh [][]string) {
	x, y := DF[d](r.Pos.X, r.Pos.Y)
	if wh[y][x] == "#" {
		return
	}
	if wh[y][x] == "." {
		wh[r.Pos.Y][r.Pos.X] = "."
		wh[y][x] = "@"
		r.Pos.Y = y
		r.Pos.X = x
		return
	}
	q := []DBox{}
	q = append(q, DBox{
		Char: wh[y][x],
		Old: Vec{
			X: x,
			Y: y,
		},
	})
	if wh[y][x] == "[" {
		q = append(q, DBox{
			Char: "]",
			Old: Vec{
				X: x + 1,
				Y: y,
			},
		})
	} else {
		q = append(q, DBox{
			Char: "[",
			Old: Vec{
				X: x - 1,
				Y: y,
			},
		})
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		x1, y1 := DF[d](curr.Old.X, curr.Old.Y)
		curr.New = Vec{
			X: x1,
			Y: y1,
		}
		switch wh[y1][x1] {
		case "#":
			return
		case "[":
			q = append(q, DBox{
				Char: "[",
				Old: Vec{
					X: x1,
					Y: y1,
				},
			})
			q = append(q, DBox{
				Char: "]",
				Old: Vec{
					X: x1 + 1,
					Y: y1,
				},
			})
		case "]":
			q = append(q, DBox{
				Char: "[",
				Old: Vec{
					X: x1 - 1,
					Y: y1,
				},
			})
			q = append(q, DBox{
				Char: "]",
				Old: Vec{
					X: x1 + 1,
					Y: y1,
				},
			})
		}

	}

}

func calcGps(wh [][]string) int {
	res := 0
	for i, row := range wh {
		for j, char := range row {
			if char == "O" {
				res += (100 * ((len(wh) - 1) - i)) + (j)
			}
		}
	}

	return res
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	warehouse, dirs := splitDirs(l)
	robot := findRobot(warehouse)
	for _, dir := range dirs {
		move(dir, robot, warehouse)
	}
	total := calcGps(warehouse)

	fmt.Println("Part one: ", total)
}

func main() {
	partOne()
}
