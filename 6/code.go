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

func partOne() {
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

	fwd = sy
	crs = sx
	edge := top
	other := right

	for fwd <= edge {
		if m[fwd][crs].Char == "#" {
			m = utils.RotateMatrix(m)
			fwd -= 1
			fwd, crs = crs, top-fwd
			edge, other = other, edge
			continue
		}
		res.Add(m[fwd][crs].Sig)
		fwd += 1
	}

	fmt.Println(res.Card())
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
	var d [][]string

	top := len(g) - 1
	right := len(g[0]) - 1
	res = NewResult()
	dirs := []string{"N", "E", "S", "W"}
	didx := 0

	for i := range g {
		m = append(m, []Cell{})
		d = append(d, make([]string, len(g[i])))
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
			d = utils.RotateMatrix(d)
			fwd -= 1
			fwd, crs = crs, top-fwd
			edge, other = other, edge
			didx = nextDir(didx)

			continue
		}
		if fwd+1 <= edge && m[fwd+1][crs].Char != "#" {
			d[fwd][crs] += dirs[didx]
		}
		for i := crs; i <= other; i++ {
			if m[fwd][i].Char == "#" {
				break
			}
			if i+1 <= other && fwd+1 <= edge && (strings.Contains(d[fwd][i], dirs[nextDir(didx)])) {
				_, ok := ps[m[fwd+1][crs].Coords]
				if !ok {
					place += 1
				}
				break
			}
		}
		res.Add(m[fwd][crs].Sig)
		fwd += 1
	}

	fmt.Println(place)

}

func main() {
	partOne()
	partTwo()
}
