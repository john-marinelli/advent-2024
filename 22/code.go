package main

import (
	"advent-2024/utils"
	"fmt"
)

type Seq struct {
	D Deltas
	P int
}

type Deltas struct {
	A int
	B int
	C int
	D int
}

func getInitial(lines []string) []int {
	init := []int{}
	for _, l := range lines {
		init = append(init, utils.StrToInt(l))
	}

	return init
}

func getNthSecret(secret int, n int) (int, []Seq) {
	curr := secret
	prices := []int{}
	deltas := []int{}
	for range n {
		prev := curr % 10
		curr = ((curr * 64) ^ curr) % 16777216
		curr = ((curr / 32) ^ curr) % 16777216
		curr = ((curr * 2048) ^ curr) % 16777216
		prices = append(prices, curr%10)
		deltas = append(deltas, (curr%10)-prev)
	}

	seqs := []Seq{}

	for i := 3; i < len(prices); i++ {
		seqs = append(seqs, Seq{
			D: Deltas{
				A: deltas[i-3],
				B: deltas[i-2],
				C: deltas[i-1],
				D: deltas[i],
			},
			P: prices[i],
		})
	}

	return curr, seqs
}

func getMaxMap(seqs []Seq) map[Deltas]int {
	m := make(map[Deltas]int)
	for _, s := range seqs {
		p, ok := m[s.D]
		if !ok {
			m[s.D] = s.P
		} else if p < s.P {
			m[s.D] = s.P
		}
	}

	return m
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	init := getInitial(l)
	total := 0
	for _, i := range init {
		t, _ := getNthSecret(i, 2000)
		total += t
	}
	fmt.Println(total)
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	init := getInitial(l)
	pMaps := []map[Deltas]int{}
	for _, i := range init {
		_, seqs := getNthSecret(i, 2000)
		pMaps = append(pMaps, getMaxMap(seqs))
	}
	mainMap := make(map[Deltas]int)
	for i := range pMaps {
		for j := range pMaps[i] {
			if _, ok := mainMap[j]; !ok {
				mainMap[j] = pMaps[i][j]
			} else {
				mainMap[j] += pMaps[i][j]
			}
		}
	}

	max := 0
	for k := range mainMap {
		if mainMap[k] > max {
			max = mainMap[k]
		}
	}

	fmt.Println(max)
}

func main() {
	partOne()
	partTwo()
}
