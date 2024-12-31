package main

import (
	"advent-2024/utils"
	"fmt"
	"strings"
)

type PathRes struct {
	Paths [][]string
}

type AdjMatrix map[string]map[string]bool

func createGraph(lines []string) AdjMatrix {
	g := make(AdjMatrix)
	for _, l := range lines {
		spl := strings.Split(l, "-")
		if _, ok := g[spl[0]]; !ok {
			g[spl[0]] = make(map[string]bool)
		}
		g[spl[0]][spl[1]] = true
		if _, ok := g[spl[1]]; !ok {
			g[spl[1]] = make(map[string]bool)
		}
		g[spl[1]][spl[0]] = true
	}
	return g
}

func dfs(s string, graph AdjMatrix, p *PathRes, path []string, l int, n int) {
	path = append(path, s)
	if l == n {
		p.Paths = append(p.Paths, path)
		return
	}

	for nd := range graph[s] {
		nw := []string{}
		nw = append(nw, path...)
		dfs(nd, graph, p, nw, l+1, n)
	}
}

func checkValid(path []string, cSet map[string]bool) bool {
	if path[0] != path[len(path)-1] {
		return false
	}
	pms := utils.GetPermutations(path[:len(path)-1])
	perms := []string{}
	for _, pm := range pms {
		perms = append(perms, strings.Join(pm, ","))
	}
	for _, pm := range perms {
		if _, ok := cSet[pm]; ok {
			return false
		}
	}

	for _, pm := range perms {
		cSet[pm] = true
	}

	return true
}

func getCycles(p *PathRes) [][]string {
	cSet := make(map[string]bool)
	cycles := [][]string{}
	for _, path := range p.Paths {
		if checkValid(path, cSet) {
			cycles = append(cycles, path[:len(path)-1])
		}
	}

	return cycles
}

func filterForChar(cycles [][]string, char string) [][]string {
	filtered := [][]string{}
	for _, c := range cycles {
		for _, n := range c {
			if strings.HasPrefix(n, char) {
				filtered = append(filtered, c)
				break
			}
		}
	}

	return filtered
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	graph := createGraph(l)
	res := &PathRes{
		Paths: [][]string{},
	}
	for node := range graph {
		dfs(node, graph, res, []string{}, 0, 3)
	}

	cycles := getCycles(res)
	filtered := filterForChar(cycles, "t")
	fmt.Println(len(filtered))
}

func main() {
	partOne()
}
