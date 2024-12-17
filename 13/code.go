package main

import (
	"advent-2024/utils"
	"fmt"
	"regexp"
)

type Res struct {
	A int
	B int
}

type Vec struct {
	X int
	Y int
}

type Buttons struct {
	A Vec
	B Vec
}

func parseInput(lines []string) ([]Vec, []Buttons) {
	targets := []Vec{}
	buttons := []Buttons{}
	dRe := regexp.MustCompile(`X=([0-9]+), Y=([0-9]+)`)
	bRe := regexp.MustCompile(`X\+([0-9]+), Y\+([0-9]+)`)

	i := 0
	for i+2 < len(lines) {
		vals := bRe.FindStringSubmatch(lines[i])
		a := Vec{
			X: utils.StrToInt(vals[1]),
			Y: utils.StrToInt(vals[2]),
		}
		vals = bRe.FindStringSubmatch(lines[i+1])
		b := Vec{
			X: utils.StrToInt(vals[1]),
			Y: utils.StrToInt(vals[2]),
		}
		buttons = append(buttons, Buttons{
			A: a,
			B: b,
		})
		dVals := dRe.FindStringSubmatch(lines[i+2])
		targets = append(targets, Vec{
			X: utils.StrToInt(dVals[1]),
			Y: utils.StrToInt(dVals[2]),
		})
		i += 4
	}

	return targets, buttons
}

func findOptimal(targets []Vec, buttons []Buttons) int {
	res := []Res{}
	for i, t := range targets {
		if (t.X%utils.Gcd(buttons[i].A.X, buttons[i].B.X) != 0) ||
			(t.Y%utils.Gcd(buttons[i].A.Y, buttons[i].B.Y) != 0) {
			continue
		}

		a := (((t.Y * buttons[i].B.X) - (t.X * buttons[i].B.Y)) /
			((buttons[i].A.Y * buttons[i].B.X) - (buttons[i].B.Y * buttons[i].A.X)))
		b := (t.X - buttons[i].A.X*a) / buttons[i].B.X

		res = append(res, Res{
			A: a,
			B: b,
		})
	}

	total := 0
	for _, r := range res {
		total += 3*r.A + r.B
	}

	return total
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	targets, buttons := parseInput(l)
	res := findOptimal(targets, buttons)
	fmt.Println(res)
}

func main() {
	partOne()
}
