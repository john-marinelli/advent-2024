package main

import (
	"advent-2024/utils"
	"fmt"
	"math"
)

type Dk struct {
	Vertices  map[Node]bool
	Graph     AdjMatrix
	Distances map[Node]int
	Paths     map[Node]bool
	Source    Node
}

type Point struct {
	X int
	Y int
}

type Node struct {
	Pos    Point
	Facing Point
}

type AdjMatrix map[Node]map[Node]int
type Visited map[Node]bool

var Dirs []Point = []Point{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

var NilNode Node = Node{}

func NewDk(v map[Node]bool, g AdjMatrix, s Node) *Dk {
	var dk Dk
	dists := make(map[Node]int)
	paths := make(map[Node]bool)
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

func (d *Dk) minDistance() Node {
	min := math.MaxInt
	var minNode Node

	for v := range d.Vertices {
		if d.Distances[v] < min && !d.Paths[v] {
			min = d.Distances[v]
			minNode = v
		}
	}

	return minNode
}

func (d *Dk) exists(x Node, y Node) bool {
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

func (d *Dk) MinCost(end Point) int {
	min := math.MaxInt
	for _, dir := range Dirs {
		cost, ok := d.Distances[Node{Pos: end, Facing: dir}]
		if ok && cost < min {
			min = cost
		}
	}

	return min
}

func findStartEnd(grid [][]string) (Point, Point) {
	start := Point{}
	end := Point{}
	for i, row := range grid {
		for j, char := range row {
			if char == "S" {
				start = Point{
					X: j,
					Y: i,
				}
			}
			if char == "E" {
				end = Point{
					X: j,
					Y: i,
				}
			}
		}
	}
	return start, end
}

func translate(p Point, d Point) Point {
	return Point{
		X: p.X + d.X,
		Y: p.Y + d.Y,
	}
}

func valid(grid [][]string, p Point) bool {
	if p.X < 0 || p.X >= len(grid[0]) || p.Y < 0 || p.Y >= len(grid) || grid[p.Y][p.X] == "#" {
		return false
	}

	return true
}

func opposite(d Point) Point {
	return Point{
		X: d.X * -1,
		Y: d.Y * -1,
	}
}

func parseCosts(
	grid [][]string,
	adj AdjMatrix,
	curr Node,
	prev Node,
	visited Visited,
) {
	if curr.Facing != prev.Facing {
		if _, ok := adj[prev]; !ok {
			adj[prev] = make(map[Node]int)
		}
		adj[prev][curr] = 1001
	} else if prev != NilNode {
		if _, ok := adj[prev]; !ok {
			adj[prev] = make(map[Node]int)
		}
		adj[prev][curr] = 1
	}

	if _, ok := visited[curr]; ok {
		return
	}

	visited[curr] = true

	for _, dir := range Dirs {
		t := translate(curr.Pos, dir)
		if valid(grid, t) {
			parseCosts(
				grid,
				adj,
				Node{
					Pos:    t,
					Facing: dir,
				},
				curr,
				visited,
			)
		}
	}
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	grid := utils.ParseCartesian(l)
	start, end := findStartEnd(grid)
	visited := make(Visited)
	adj := make(AdjMatrix)
	source := Node{Pos: start, Facing: Dirs[2]}

	parseCosts(grid, adj, source, NilNode, visited)

	dk := NewDk(visited, adj, source)
	dk.ShortestPath()
	fmt.Println(dk.MinCost(end))
}

func main() {
	partOne()
}
