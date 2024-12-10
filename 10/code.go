package main

import (
	"advent-2024/utils"
	"fmt"
	"slices"
	"strings"
)

// Deltas is [N, E, S, W]

type DirFunc func(x int, y int) (int, int)

var ops map[string]DirFunc = map[string]DirFunc{
	"N": func(x int, y int) (int, int) { return x, y + 1 },
	"S": func(x int, y int) (int, int) { return x, y - 1 },
	"E": func(x int, y int) (int, int) { return x + 1, y },
	"W": func(x int, y int) (int, int) { return x - 1, y },
}

type Point struct {
	X      int
	Y      int
	Height int
	Deltas map[string]int
}

func dfs(
	adj map[*Point][]*Point,
	visited map[*Point]bool,
	point *Point,
	length int,
	res *int,
) {
	visited[point] = true

	if length == 9 {
		*res += 1
		return
	}

	for i := range adj[point] {
		if !visited[adj[point][i]] {
			dfs(adj, visited, adj[point][i], length+1, res)
		}
	}
}

func dfsPaths(
	adj map[*Point][]*Point,
	path []*Point,
	length int,
	paths [][]*Point,
) [][]*Point {
	curr := path[len(path)-1]
	if _, ok := adj[curr]; ok {
		for i := range adj[curr] {
			newPath := append(path, adj[curr][i])
			paths = dfsPaths(adj, newPath, length+1, paths)
		}
	} else if curr.Height == 9 {
		paths = append(paths, path)
	}
	return paths

}

func search(adj map[*Point][]*Point, source *Point, allUnique bool) int {
	score := new(int)
	*score = 0
	visited := make(map[*Point]bool)
	for k := range adj {
		visited[k] = false
	}

	if allUnique {
		paths := [][]*Point{}
		paths = dfsPaths(adj, []*Point{source}, 0, paths)
		return len(paths)
	} else {
		dfs(adj, visited, source, 0, score)
	}

	return *score
}

func createAdjList(grid [][]*Point) map[*Point][]*Point {
	adj := make(map[*Point][]*Point)
	for i := range grid {
		for j := range grid[i] {
			for k := range grid[i][j].Deltas {
				d := grid[i][j].Deltas[k]
				if d == 1 {
					x, y := ops[k](grid[i][j].X, grid[i][j].Y)
					if _, ok := adj[grid[i][j]]; !ok {
						adj[grid[i][j]] = []*Point{}
					}
					adj[grid[i][j]] = append(adj[grid[i][j]], grid[y][x])
				}
			}
		}
	}

	return adj
}

func parseGrid(l []string) [][]*Point {
	res := [][]*Point{}

	slices.Reverse(l)
	for i := range l {
		res = append(res, []*Point{})
		row := strings.Split(l[i], "")
		for j := range row {
			res[i] = append(res[i], &Point{
				X:      j,
				Y:      i,
				Height: utils.StrToInt(row[j]),
			})
		}
	}

	return res
}

func calcDelta(grid [][]*Point, p *Point, dx int, dy int) int {
	return grid[p.Y+dy][p.X+dx].Height - p.Height
}

func fillDeltas(grid [][]*Point) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j].Deltas = make(map[string]int)
			if j+1 < len(grid[i]) {
				grid[i][j].Deltas["E"] = calcDelta(grid, grid[i][j], 1, 0)
			}
			if j-1 >= 0 {
				grid[i][j].Deltas["W"] = calcDelta(grid, grid[i][j], -1, 0)
			}
			if i+1 < len(grid) {
				grid[i][j].Deltas["N"] = calcDelta(grid, grid[i][j], 0, 1)
			}
			if i-1 >= 0 {
				grid[i][j].Deltas["S"] = calcDelta(grid, grid[i][j], 0, -1)
			}
		}
	}
}

func prep() map[*Point][]*Point {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	grid := parseGrid(l)
	fillDeltas(grid)
	adj := createAdjList(grid)

	return adj
}

func findTrails(partTwo bool) int {
	adj := prep()
	res := 0
	for p := range adj {
		if p.Height != 0 {
			continue
		}
		res += search(adj, p, partTwo)
	}

	return res
}

func partOne() {
	fmt.Println("Part one: ", findTrails(false))
}

func partTwo() {
	fmt.Println("Part two: ", findTrails(true))
}

func main() {
	partOne()
	partTwo()
}
