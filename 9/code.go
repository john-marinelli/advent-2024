package main

import (
	"advent-2024/utils"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Block struct {
	Size    int
	File    bool
	OrigIdx int
}

type Storage struct {
	Full []*Block
	free *Block
	file *Block
}

func (s *Storage) consumeFree(freeIdx int, fileIdx int) {
	leftover := s.free.Size - s.file.Size
	if leftover > 0 {
		s.free.Size = leftover
		s.Full[fileIdx] = &Block{
			File: false,
			Size: s.file.Size,
		}
		s.Full = slices.Insert(s.Full, freeIdx, s.file)
		return
	}

	s.Full[freeIdx], s.Full[fileIdx] = s.Full[fileIdx], s.Full[freeIdx]
}

func (s *Storage) compactFile(fIdx int) {
	for i := range fIdx {
		curr := s.Full[i]
		if !curr.File && curr.Size >= s.file.Size {
			s.free = curr
			s.consumeFree(i, fIdx)
			break
		}
	}
}

func (s *Storage) Compact() {
	j := len(s.Full) - 1

	for j > 0 {
		if s.Full[j].File {
			s.file = s.Full[j]
			s.compactFile(j)
		}
		j -= 1
	}
}

func (s *Storage) Render() []int {
	res := []int{}
	for i := range s.Full {
		if s.Full[i].File {
			for range s.Full[i].Size {
				res = append(res, s.Full[i].OrigIdx)
			}
		} else {
			for range s.Full[i].Size {
				res = append(res, -1)
			}
		}
	}

	return res
}

func NewStorage(preProc []string) *Storage {
	mem := []*Block{}
	cnt := 0
	for i, p := range preProc {
		if i%2 == 0 {
			mem = append(mem, &Block{
				OrigIdx: cnt,
				Size:    utils.StrToInt(p),
				File:    true,
			})
			cnt += 1
		} else {
			mem = append(mem, &Block{
				Size: utils.StrToInt(p),
				File: false,
			})
		}
	}

	return &Storage{
		Full: mem,
	}
}

func processMem(preProc []string) []int {
	res := []int{}
	blockIdx := 0
	for i, p := range preProc {
		if i%2 == 0 {
			size := utils.StrToInt(p)
			for range size {
				res = append(res, blockIdx)
			}

			blockIdx += 1
		} else {
			size := utils.StrToInt(p)
			for range size {
				res = append(res, -1)
			}
		}
	}

	return res
}

func compact(mem []int) []int {
	i := 0
	j := len(mem) - 1

	for i < len(mem) && j >= 0 && j > i {
		if mem[i] < 0 && mem[j] >= 0 {
			mem[i], mem[j] = mem[j], mem[i]
			i += 1
			j -= 1
			continue
		}
		if mem[i] >= 0 {
			i += 1
		}
		if mem[j] < 0 {
			j -= 1
		}
	}

	return mem
}

func checksum(mem []int) int {
	res := 0
	for i, m := range mem {
		if m >= 0 {
			res += i * m
		}
	}

	return res
}

func partOne() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	preProc := strings.Split(l[0], "")
	mem := processMem(preProc)
	mem = compact(mem)
	res := checksum(mem)
	fmt.Println("Part one: ", res)
}

func partTwo() {
	l, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	preProc := strings.Split(l[0], "")
	strg := NewStorage(preProc)
	strg.Compact()
	res := checksum(strg.Render())
	fmt.Println("Part two: ", res)
}

func main() {
	// partOne()
	start := time.Now()
	partTwo()
	fmt.Printf("Ex time: %v", time.Since(start))
}
