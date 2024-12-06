package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	var o []string

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		o = append(o, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return o, nil
}

func SplitNumericalLines(fLines []string) ([][]int, error) {
	var o [][]int

	for _, l := range fLines {
		var inter []int
		for _, k := range strings.Fields(l) {
			conv, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			inter = append(inter, conv)
		}
		o = append(o, inter)

	}

	return o, nil
}

func SplitLines(fLines []string) ([][]string, error) {
	var o [][]string

	for _, l := range fLines {
		o = append(o, strings.Split(l, ""))
	}

	return o, nil
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func RemoveEle(s []int, i int) []int {
	var result []int

	for idx := range s {
		if idx == i {
			continue
		}
		result = append(result, s[idx])
	}

	return result
}

func GetFileString(path string) string {
	s, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(s)
}

func MatchReExact(str string, re string, em int) [][]string {
	var result [][]string
	r := regexp.MustCompile(re)

	res := r.FindAllStringSubmatch(str, -1)

	for _, m := range res {
		var match []string
		if len(m) != (em + 1) {
			continue
		}
		for i := 1; i < len(m); i++ {

			match = append(match, m[i])
		}
		result = append(result, match)
	}
	return result
}

func MatchRe(str string, re string) [][]string {
	r := regexp.MustCompile(re)

	res := r.FindAllStringSubmatch(str, -1)

	return res
}

func SplitPage(page []string) ([]string, []string) {
	sep := 0
	for i := range page {
		if page[i] == "" {
			sep = i
			break
		}
	}

	return page[:sep], page[sep+1:]
}

func IsIn[T comparable](ele T, m map[T]bool) bool {
	_, ok := m[ele]
	return ok
}
