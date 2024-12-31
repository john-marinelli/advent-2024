package main

import (
	"advent-2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	Left  *Wire
	Right *Wire
	Out   *Wire
	Op    string
}

type Wire struct {
	Name  string
	State int
	Set   bool
}

type LogicFunc func(a int, b int) int

var Ops map[string]LogicFunc = map[string]LogicFunc{
	"AND": func(a int, b int) int {
		if a == 1 && b == 1 {
			return 1
		}
		return 0
	},
	"OR": func(a int, b int) int {
		if a == 1 || b == 1 {
			return 1
		}
		return 0
	},
	"XOR": func(a int, b int) int {
		if (a == 1 && b == 0) || (a == 0 && b == 1) {
			return 1
		}
		return 0
	},
}

func createUnsetWire(name string, wires map[string]*Wire) {
	if _, ok := wires[name]; !ok {
		wires[name] = &Wire{
			Name: name,
			Set:  false,
		}
	}
}

func parse(lines []string) (map[string]*Wire, [][]string) {
	isGate := false
	wires := make(map[string]*Wire)
	gateList := [][]string{}
	for _, line := range lines {
		if line == "" {
			isGate = true
			continue
		}
		if !isGate {
			spl := strings.Split(line, ":")
			wires[spl[0]] = &Wire{
				Name:  spl[0],
				State: utils.StrToInt(strings.TrimSpace(spl[1])),
				Set:   true,
			}
			continue
		}
		spl := strings.Fields(line)
		gateList = append(gateList, []string{spl[0], spl[1], spl[2], spl[4]})
		createUnsetWire(spl[0], wires)
		createUnsetWire(spl[2], wires)
		createUnsetWire(spl[4], wires)
	}

	return wires, gateList
}

func assembleGates(wires map[string]*Wire, gateList [][]string) []*Gate {
	gates := []*Gate{}
	for _, gate := range gateList {
		gates = append(gates, &Gate{
			Left:  wires[gate[0]],
			Right: wires[gate[2]],
			Op:    gate[1],
			Out:   wires[gate[3]],
		})
	}

	return gates
}

func fillQ(q []*Gate, inQ map[*Gate]bool, gates []*Gate) []*Gate {
	for _, g := range gates {
		_, ok := inQ[g]
		if g.Left.Set && g.Right.Set && !g.Out.Set && !ok {
			q = append(q, g)
			inQ[g] = true
		}
	}

	return q
}

func runLogic(gates []*Gate) {
	q := []*Gate{}
	inQ := make(map[*Gate]bool)
	q = fillQ(q, inQ, gates)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		curr.Out.State = Ops[curr.Op](curr.Left.State, curr.Right.State)
		curr.Out.Set = true
		q = fillQ(q, inQ, gates)
	}
}

func getZResult(wires map[string]*Wire) int {
	type Z struct {
		Order int
		State string
	}
	z := []Z{}
	for name, wire := range wires {
		if strings.HasPrefix(name, "z") {
			z = append(z, Z{
				Order: utils.StrToInt(strings.ReplaceAll(name, "z", "")),
				State: strconv.Itoa(wire.State),
			})
		}
	}
	sort.Slice(z, func(i int, j int) bool {
		return z[i].Order > z[j].Order
	})

	binStr := []string{}
	for _, entry := range z {
		binStr = append(binStr, entry.State)
	}

	res, err := strconv.ParseInt(strings.Join(binStr, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(res)
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	wires, gateList := parse(l)
	gates := assembleGates(wires, gateList)
	runLogic(gates)
	res := getZResult(wires)
	fmt.Println(res)
}

func main() {
	partOne()
}
