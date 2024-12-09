package main

import (
	"advent-2024/utils"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Slope struct {
	Rise int
	Run  int
}

func calcGcd(a int, b int) int {
	if a == 0 {
		return b
	}

	return calcGcd(b%a, a)
}

func collectTowers(city [][]string) map[string][]Point {
	towers := make(map[string][]Point)
	for y, row := range city {
		for x, tower := range row {
			if tower != "." {
				if _, ok := towers[tower]; !ok {
					towers[tower] = []Point{}
				}
				towers[tower] = append(
					towers[tower],
					Point{
						X: x,
						Y: y,
					},
				)
			}
		}
	}

	return towers
}

func slopes(a Point, b Point) (Slope, Slope) {
	return Slope{
			Rise: a.Y - b.Y,
			Run:  a.X - b.X,
		},
		Slope{
			Rise: b.Y - a.Y,
			Run:  b.X - a.X,
		}
}

func calcAntinodes(tows []Point, maxX int, maxY int, res map[Point]bool) {
	for i, a := range tows {
		for j, b := range tows {
			if i != j {
				slopeA, slopeB := slopes(a, b)
				an1 := Point{
					X: a.X + slopeA.Run,
					Y: a.Y + slopeA.Rise,
				}
				an2 := Point{
					X: b.X + slopeB.Run,
					Y: b.Y + slopeB.Rise,
				}

				if an1.X < maxX && an1.Y < maxY && an1.X >= 0 && an1.Y >= 0 {
					res[an1] = true
				}
				if an2.X < maxX && an2.Y < maxY && an2.X >= 0 && an2.Y >= 0 {
					res[an2] = true
				}
			}
		}
	}
}

func calcAntinodeLines(tows []Point, maxX int, maxY int, res map[Point]bool) {
	for i, a := range tows {
		for j, b := range tows {
			if i != j {
				slope, _ := slopes(a, b)
				getLine(slope, a.X, a.Y, maxX, maxY, res)
			}
		}
	}
}

func getLine(slope Slope, x int, y int, maxX int, maxY int, res map[Point]bool) {
	gcd := calcGcd(slope.Rise, slope.Run)
	deltaX := slope.Run / gcd
	deltaY := slope.Rise / gcd

	currX := x
	currY := y
	for currX < maxX && currY < maxY && currX >= 0 && currY >= 0 {
		res[Point{X: currX, Y: currY}] = true

		currX += deltaX
		currY += deltaY
	}

	currX = x
	currY = y

	for currX < maxX && currY < maxY && currX >= 0 && currY >= 0 {

		res[Point{X: currX, Y: currY}] = true

		currX -= deltaX
		currY -= deltaY
	}
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	city := utils.ParseCartesian(l)
	towers := collectTowers(city)
	res := make(map[Point]bool)
	maxY := len(city)
	maxX := len(city[0])

	for towerType := range towers {
		calcAntinodes(towers[towerType], maxX, maxY, res)
	}

	fmt.Println(len(res))
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	city := utils.ParseCartesian(l)
	towers := collectTowers(city)
	res := make(map[Point]bool)
	maxY := len(city)
	maxX := len(city[0])

	for towerType := range towers {
		calcAntinodeLines(towers[towerType], maxX, maxY, res)
	}

	fmt.Println(len(res))
}

func main() {
	partOne()
	partTwo()
}
