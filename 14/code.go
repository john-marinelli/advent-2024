package main

import (
	"advent-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Vec struct {
	X int
	Y int
}

type Robot struct {
	Init Vec
	Vel  Vec
	Curr Vec
}

func parseRobots(lines []string) []*Robot {
	robots := []*Robot{}
	re := regexp.MustCompile(`-?[0-9]+`)
	for _, line := range lines {
		nums := re.FindAllStringSubmatch(line, -1)
		robots = append(robots, &Robot{
			Init: Vec{
				X: utils.StrToInt(nums[0][0]),
				Y: utils.StrToInt(nums[1][0]),
			},
			Vel: Vec{
				X: utils.StrToInt(nums[2][0]),
				Y: utils.StrToInt(nums[3][0]),
			},
			Curr: Vec{
				X: utils.StrToInt(nums[0][0]),
				Y: utils.StrToInt(nums[1][0]),
			},
		})
	}

	return robots
}

func step(robot *Robot, x int, y int) {
	robot.Curr.X = utils.PyMod((robot.Curr.X + robot.Vel.X), x)
	robot.Curr.Y = utils.PyMod((robot.Curr.Y + robot.Vel.Y), y)
}

func allAtInit(robots []*Robot) bool {
	for _, r := range robots {
		if r.Curr != r.Init {
			return false
		}
	}

	return true
}

func iterate(robots []*Robot, x int, y int, iters int) []*Robot {
	for _ = range iters {
		for _, r := range robots {
			step(r, x, y)
		}
	}
	return robots
}

func count(robots []*Robot, x int, y int) int {
	yh := (y / 2)
	xh := (x / 2)
	fmt.Println(xh, " ", yh)
	ll := 0
	lr := 0
	ul := 0
	ur := 0

	for _, r := range robots {
		if r.Curr.X < xh && r.Curr.Y < yh {
			ll += 1
		} else if r.Curr.X < xh && r.Curr.Y > yh {
			ul += 1
		} else if r.Curr.X > xh && r.Curr.Y < yh {
			lr += 1
		} else if r.Curr.X > xh && r.Curr.Y > yh {
			ur += 1
		}
	}

	return ll * lr * ul * ur
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	wide := 101
	tall := 103
	robots := parseRobots(l)
	iterate(robots, wide, tall, 100)
	res := count(robots, wide, tall)
	fmt.Println("Part one: ", res)
}

func createDisplay(x int, y int) [][]string {
	out := [][]string{}
	for range y {
		out = append(out, []string{})
		for range x {
			out[len(out)-1] = append(out[len(out)-1], ".")
		}
	}

	return out
}

func displayRobots(robots []*Robot, x int, y int) {
	display := createDisplay(x, y)
	for _, robot := range robots {
		if display[robot.Curr.Y][robot.Curr.X] == "." {
			display[robot.Curr.Y][robot.Curr.X] = "1"
		} else {
			display[robot.Curr.Y][robot.Curr.X] = strconv.Itoa(
				utils.StrToInt(display[robot.Curr.Y][robot.Curr.X]) + 1,
			)
		}
	}
	for i := range display {
		fmt.Println(display[i])
	}
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	wide := 101
	tall := 103
	robots := parseRobots(l)
	iterate(robots, wide, tall, 1000)
	res := count(robots, wide, tall)
	fmt.Println("Part one: ", res)
}

func main() {
	partOne()
	partTwo()
}
