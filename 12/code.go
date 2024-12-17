package main

import (
	"advent-2024/utils"
	"fmt"
)

type Square struct {
	Type  string
	Sides []Side
}

type Point struct {
	X int
	Y int
}

type Plot struct {
	Squares []Square
	Size    int
}

type Side struct {
	Begin Point
	End   Point
}

func exploreGarden(exp [][]Square) []Plot {
	visited := make(map[Point]bool)

	plots := []Plot{}

	for i, row := range exp {
		for j, square := range row {
			q := []Point{}
			q = append(q, Point{X: j, Y: i})
			if _, ok := visited[q[0]]; ok {
				continue
			}
			plots = append(plots, Plot{})
			plots[len(plots)-1].Size = 0
			for len(q) > 0 {
				curr := q[0]
				q = q[1:]
				if _, ok := visited[curr]; ok {
					continue
				}
				plots[len(plots)-1].Size += 1
				visited[curr] = true
				plots[len(plots)-1].Squares = append(plots[len(plots)-1].Squares, exp[curr.Y][curr.X])
				if curr.Y+1 < len(exp) && exp[curr.Y+1][curr.X].Type == square.Type {
					q = append(q, Point{
						X: curr.X,
						Y: curr.Y + 1,
					})
				}
				if curr.X+1 < len(row) && exp[curr.Y][curr.X+1].Type == square.Type {
					q = append(q, Point{
						X: curr.X + 1,
						Y: curr.Y,
					})
				}
				if curr.X-1 >= 0 && exp[curr.Y][curr.X-1].Type == square.Type {
					q = append(q, Point{
						X: curr.X - 1,
						Y: curr.Y,
					})
				}
				if curr.Y-1 >= 0 && exp[curr.Y-1][curr.X].Type == square.Type {
					q = append(q, Point{
						X: curr.X,
						Y: curr.Y - 1,
					})
				}
			}
		}
	}

	return plots
}

func addSide(p Point, sx int, sy int, ex int, ey int) Side {
	s := Side{
		Begin: Point{
			X: p.X + sx,
			Y: p.Y + sy,
		},
		End: Point{
			X: p.X + ex,
			Y: p.Y + ey,
		},
	}
	return s
}

func getGardenCoordinates(grid [][]string) [][]Square {
	o := [][]Square{}
	for range len(grid) {
		o = append(o, []Square{})
	}
	for i, row := range grid {
		for j, square := range row {
			o[j] = append(o[j], Square{
				Type: square,
				Sides: []Side{
					{Begin: Point{X: j, Y: i}, End: Point{X: j + 1, Y: i}},
					{Begin: Point{X: j + 1, Y: i}, End: Point{X: j + 1, Y: i + 1}},
					{Begin: Point{X: j + 1, Y: i + 1}, End: Point{X: j, Y: i + 1}},
					{Begin: Point{X: j, Y: i + 1}, End: Point{X: j, Y: i}},
				},
			})
		}
	}

	return o
}

func calculateCost(plot Plot) int {
	res := 0
	sides := make(map[Side]bool)

	for _, s := range plot.Squares {
		for _, side := range s.Sides {
			sides[side] = true
		}
	}

	for _, s := range plot.Squares {
		for _, side := range s.Sides {
			_, fok := sides[side]
			_, bok := sides[Side{Begin: side.End, End: side.Begin}]
			if fok && bok {
				continue
			}
			res += 1
		}
	}

	return plot.Size * res
}

func checkExterior(a Side, set map[Side]bool) bool {
	_, bok := set[Side{Begin: a.End, End: a.Begin}]
	return !bok
}

func perpendicular(a Side, b Side) bool {
	if a.Begin.X == a.End.X {
		return a.Begin.X != b.End.X
	}

	return a.Begin.Y != b.End.Y
}

func calculateDiscounted(plot Plot) int {
	sides := make(map[Side]bool)
	connector := make(map[Point][]Side)
	checked := make(map[Side]bool)

	for _, s := range plot.Squares {
		for _, side := range s.Sides {
			sides[side] = true
			connector[side.Begin] = append(connector[side.Begin], side)
		}
	}

	corners := 0

	for side := range sides {
		if !checkExterior(side, sides) {
			continue
		}

		for _, ext := range connector[side.End] {
			if !checkExterior(ext, sides) {
				continue
			}
			if perpendicular(side, ext) {
				if _, ok := checked[ext]; ok {
					continue
				}
				corners += 1
				checked[ext] = true
			}
		}
	}

	return corners * plot.Size
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	grid := utils.ParseCartesian(l)
	exp := getGardenCoordinates(grid)
	explored := exploreGarden(exp)
	total := 0

	for _, plot := range explored {
		total += calculateCost(plot)
	}

	fmt.Println("Part one: ", total)

}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	grid := utils.ParseCartesian(l)
	exp := getGardenCoordinates(grid)
	explored := exploreGarden(exp)

	total := 0
	for _, e := range explored {
		total += calculateDiscounted(e)
	}

	fmt.Println("Part two: ", total)

}

func main() {
	partOne()
	partTwo()
}
