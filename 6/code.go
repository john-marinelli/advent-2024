package main

import (
	"advent-2024/utils"
	"fmt"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Cell struct {
	Sig    string
	Char   string
	Coords Coord
}

type Result struct {
	Set map[string]bool
}

func (r *Result) Add(s string) {
	r.Set[s] = true
}

func (r *Result) Card() int {
	return len(r.Set)
}

func (r *Result) In(s string) bool {
	_, ok := r.Set[s]
	return ok
}

func NewResult() *Result {
	n := &Result{}
	n.Set = make(map[string]bool)
	return n
}

func nextDir(di int) int {
	return (di + 1) % 4
}

func findLoop(
	m [][]Cell,
	sx int,
	sy int,
	top int,
	right int,
	res *Result,
	vis map[string]string,
	didx int,
	dirs []string,
) (int, bool) {

	fwd := sy
	crs := sx
	edge := top
	other := right
	backup := false

	for fwd <= edge {
		if m[fwd][crs].Char == "#" {
			//This adds an extra factor of n to time complexity
			//But it makes it easier to type this out so no one
			//can stop me
			m = utils.RotateMatrix(m)
			fwd -= 1
			backup = true
			fwd, crs = crs, top-fwd
			edge, other = other, edge
			didx = nextDir(didx)
			continue
		}
		res.Add(m[fwd][crs].Sig)
		if strings.Contains(vis[m[fwd][crs].Sig], dirs[didx]) {
			if !backup {
				return res.Card(), true
			}
		}
		vis[m[fwd][crs].Sig] += dirs[didx]

		fwd += 1
		backup = false
	}

	return res.Card(), false
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	g := utils.ParseCartesian(l)
	var m [][]Cell
	var sx int
	var sy int
	var res *Result
	vis := make(map[string]string)
	dirs := []string{"N", "E", "S", "W"}
	didx := 0

	top := len(g) - 1
	right := len(g[0]) - 1
	res = NewResult()

	for i := range g {
		m = append(m, []Cell{})
		for j := range g[i] {
			if g[i][j] == "^" {
				sx = j
				sy = i
			}
			m[i] = append(m[i], Cell{
				Sig:  fmt.Sprintf("%d,%d", j, i),
				Char: g[i][j],
			})
		}
	}

	unq, loop := findLoop(m, sx, sy, top, right, res, vis, didx, dirs)
	fmt.Println("Unique visited: ", unq, loop)
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	g := utils.ParseCartesian(l)
	var m [][]Cell
	var sx int
	var sy int
	var fwd int
	var crs int
	var res *Result

	top := len(g) - 1
	right := len(g[0]) - 1
	res = NewResult()
	vis := make(map[string]string)
	didx := 0
	dirs := []string{"N", "E", "S", "W"}

	for i := range g {
		m = append(m, []Cell{})
		for j := range g[i] {
			if g[i][j] == "^" {
				sx = j
				sy = i
			}
			m[i] = append(m[i], Cell{
				Sig:  fmt.Sprintf("%d,%d", j, i),
				Char: g[i][j],
				Coords: Coord{
					X: j,
					Y: i,
				},
			})
		}
	}

	fwd = sy
	crs = sx
	edge := top
	other := right
	place := 0
	ps := make(map[Coord]bool)

	for fwd <= edge {
		if m[fwd][crs].Char == "#" {
			m = utils.RotateMatrix(m)
			fwd -= 1
			fwd, crs = crs, top-fwd
			edge, other = other, edge
			didx = nextDir(didx)
			continue
		}

		vis[m[fwd][crs].Sig] += dirs[didx]

		if fwd+1 <= edge && m[fwd+1][crs].Char != "#" && !utils.IsIn(m[fwd+1][crs].Coords, ps) {
			_, loop := findLoop(
				utils.RotateMatrix(m),
				top-fwd,
				crs,
				other,
				edge,
				res,
				utils.CopyMap(vis),
				nextDir(didx),
				dirs,
			)
			if loop {
				place += 1
				ps[m[fwd+1][crs].Coords] = true
				g[m[fwd+1][crs].Coords.Y][m[fwd+1][crs].Coords.X] = "0"
			}
		}
		fwd += 1
	}

	fmt.Println("Obstacle placements: ", place)

}

func main() {
	partOne()
	partTwo()
}
