package main

import (
	"advent-2024/utils"
	"fmt"
	"math"
	"strings"
)

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
}

type DubList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (l *DubList) Add(value int) {
	new := &Node{Value: value}
	if l.Tail == nil {
		l.Head = new
		l.Tail = new
	} else {
		new.Prev = l.Tail
		l.Tail.Next = new
		l.Tail = new
	}
}

func intLen(a int) int {
	return int(math.Floor(math.Log10(float64(a)))) + 1
}

func split(list *DubList, node *Node) (*DubList, *Node) {
	length := intLen(node.Value)
	left := node.Value / int(math.Pow(10, float64(length/2)))
	right := node.Value - left*int(math.Pow(10, float64(length/2)))

	leftNode := &Node{
		Value: left,
	}
	rightNode := &Node{
		Value: right,
	}

	leftNode.Prev = node.Prev
	rightNode.Next = node.Next
	rightNode.Prev = leftNode
	leftNode.Next = rightNode
	if node.Prev != nil {
		node.Prev.Next = leftNode
	} else {
		list.Head = leftNode
	}
	if node.Next != nil {
		node.Next.Prev = rightNode
	} else {
		list.Tail = rightNode
	}

	list.Len += 1

	return list, rightNode
}

func parseInput(line string) *DubList {
	nums := strings.Fields(line)
	fmt.Println(nums)
	list := &DubList{}
	list.Len = 0

	for _, num := range nums {
		conv := utils.StrToInt(num)
		list.Add(conv)
		list.Len += 1
	}

	return list
}

func loadInput(line string) {

}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	list := parseInput(l[0])

	for range 75 {
		curr := list.Head
		for curr != nil {
			if curr.Value == 0 {
				curr.Value = 1
			} else if intLen(curr.Value)%2 == 0 {
				list, curr = split(list, curr)
			} else {
				curr.Value *= 2024
			}
			curr = curr.Next
		}
	}

	fmt.Println(list.Len)
}

func main() {
	partOne()
}
