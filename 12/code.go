package main

import (
	"advent-2024/utils"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Plot struct {
	Points []Point
	Sides  []Side
	Size   int
}

type Side struct {
	Begin Point
	End   Point
}

func exploreGarden(exp [][]string) []Plot {
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
				plots[len(plots)-1].Points = append(plots[len(plots)-1].Points, curr)
				plots[len(plots)-1].Size += 1
				visited[curr] = true
				if curr.Y+1 < len(exp) && exp[curr.Y+1][curr.X] == square {
					q = append(q, Point{
						X: curr.X,
						Y: curr.Y + 1,
					})
				}
				if curr.X+1 < len(row) && exp[curr.Y][curr.X+1] == square {
					q = append(q, Point{
						X: curr.X + 1,
						Y: curr.Y,
					})
				}
				if curr.X-1 >= 0 && exp[curr.Y][curr.X-1] == square {
					q = append(q, Point{
						X: curr.X - 1,
						Y: curr.Y,
					})
				}
				if curr.Y-1 >= 0 && exp[curr.Y-1][curr.X] == square {
					q = append(q, Point{
						X: curr.X,
						Y: curr.Y - 1,
					})
				}
			}
		}
	}

	for i := range plots {
		plots[i].Size /= 4
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

func searchGarden(squares []Point, dirs [][]int) map[Side]bool {
	sSet := make(map[Side]bool)
	pSet := make(map[Point]bool)

	for _, p := range squares {
		pSet[p] = true
	}

	for _, p := range squares {
		_, yOk := pSet[Point{X: p.X, Y: p.Y + 1}]
		_, xOk := pSet[Point{X: p.X + 1, Y: p.Y}]
		if !xOk || !yOk {
			continue
		}
		for _, dir := range dirs {
			s := addSide(p, dir[0], dir[1], dir[2], dir[3])
			sSet[s] = true
		}
	}

	return sSet
}

func expandGarden(grid [][]string) [][]string {
	o := [][]string{}
	for range len(grid) * 2 {
		o = append(o, []string{})
	}
	j := 0
	for _, row := range grid {
		for _, square := range row {
			fmt.Println(square)
			o[j] = append(o[j], square)
			o[j] = append(o[j], square)
			o[j+1] = append(o[j+1], square)
			o[j+1] = append(o[j+1], square)
		}
		j += 2
	}

	return o
}

func calculateCost(sides map[Side]bool, area int) int {
	res := 0

	for s := range sides {
		if _, ok := sides[Side{Begin: s.End, End: s.Begin}]; !ok {
			res += 1
		}
	}

	return area * (res + 4) / 2
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	grid := utils.ParseCartesian(l)
	exp := expandGarden(grid)
	explored := exploreGarden(exp)

	for i := range explored {
		for j := range explored[i].Points {
			fmt.Print(exp[explored[i].Points[j].Y][explored[i].Points[j].X])
		}
		fmt.Print(explored[i].Size)
		fmt.Print("\n")
	}

	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(grid[i])
	}

	fmt.Println("---------------")

	for i := len(exp) - 1; i >= 0; i-- {
		fmt.Println(exp[i])
	}

	dirs := [][]int{
		{0, 0, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
	}
	cost := 0
	for i := range explored {
		currPlot := searchGarden(explored[i].Points, dirs)
		cost += calculateCost(currPlot, explored[i].Size)
	}

	fmt.Println(cost)

}

func main() {
	partOne()
}
