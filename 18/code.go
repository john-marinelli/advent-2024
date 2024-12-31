package main

import (
	"advent-2024/utils"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	X int
	Y int
}

var Dirs []Point = []Point{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

type AdjMatrix map[Point]map[Point]int

type Dk struct {
	Graph     AdjMatrix
	Vertices  map[Point]bool
	Paths     map[Point]bool
	Distances map[Point]int
	Source    Point
}

func NewDk(v map[Point]bool, g AdjMatrix, s Point) *Dk {
	var dk Dk
	dists := make(map[Point]int)
	paths := make(map[Point]bool)
	dk.Vertices = v
	dk.Graph = g
	dk.Source = s

	for n := range v {
		dists[n] = math.MaxInt
		paths[n] = false
	}
	dk.Distances = dists
	dk.Paths = paths

	return &dk
}

func (d *Dk) minDistance() Point {
	min := math.MaxInt
	var minNode Point

	for v := range d.Vertices {
		if d.Distances[v] < min && !d.Paths[v] {
			min = d.Distances[v]
			minNode = v
		}
	}

	return minNode
}

func (d *Dk) exists(x Point, y Point) bool {
	_, ok := d.Graph[x][y]

	return ok
}

func (d *Dk) ShortestPath() {
	d.Distances[d.Source] = 0

	for range d.Vertices {
		x := d.minDistance()
		d.Paths[x] = true
		for y := range d.Vertices {
			if d.exists(x, y) && !d.Paths[y] && d.Distances[y] > (d.Distances[x]+d.Graph[x][y]) {
				d.Distances[y] = d.Distances[x] + d.Graph[x][y]
			}
		}
	}
}

func getMatrix(lines []string, width int, height int, numBytes int) (AdjMatrix, map[Point]bool, string) {
	blocked := make(map[Point]bool)
	vertices := make(map[Point]bool)

	graph := make(AdjMatrix)
	i := 0
	for i < numBytes {
		spl := strings.Split(lines[i], ",")
		blocked[Point{
			X: utils.StrToInt(spl[0]),
			Y: utils.StrToInt(spl[1]),
		}] = true
		i += 1
	}

	lastByte := lines[i-1]

	for i := range height + 1 {
		for j := range width + 1 {
			p := Point{X: j, Y: i}
			if _, ok := blocked[p]; ok {
				continue
			}
			vertices[p] = true
			for _, dir := range Dirs {
				t := Point{X: p.X + dir.X, Y: p.Y + dir.Y}
				if t.X > width || t.X < 0 || t.Y > height || t.Y < 0 {
					continue
				}
				if _, ok := blocked[t]; ok {
					continue
				}
				if _, ok := graph[p]; !ok {
					graph[p] = make(map[Point]int)
				}
				graph[p][t] = 1
			}
		}
	}

	return graph, vertices, lastByte
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	graph, vertices, _ := getMatrix(l, 70, 70, 1024)
	dk := NewDk(vertices, graph, Point{X: 0, Y: 0})
	dk.ShortestPath()
	fmt.Println(dk.Distances[Point{X: 70, Y: 70}])
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	for i := 1024; i < len(l); i++ {
		graph, vertices, lb := getMatrix(l, 70, 70, i)
		dk := NewDk(vertices, graph, Point{X: 0, Y: 0})
		dk.ShortestPath()
		if dk.Distances[Point{X: 70, Y: 70}] >= math.MaxInt {
			fmt.Println(lb)
			break
		}
	}

}

func main() {
	// partOne()
	partTwo()
}
