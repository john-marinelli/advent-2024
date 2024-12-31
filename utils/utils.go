package utils

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func CopyReverse[T any](s []T) []T {
	o := make([]T, len(s))
	copy(o, s)
	slices.Reverse(o)
	return o
}

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

func ParseCartesian(lines []string) [][]string {
	var o [][]string
	slices.Reverse(lines)
	for _, l := range lines {
		o = append(o, strings.Split(l, ""))
	}

	return o
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

func GetRange(start int, end int) []int {
	var o []int
	for i := start; i < end; i++ {
		o = append(o, i)
	}

	return o
}

func RotateMatrix[T any](in [][]T) [][]T {
	var out [][]T

	k := 0
	for j := range in[0] {
		out = append(out, []T{})
		for i := len(in) - 1; i >= 0; i-- {
			out[k] = append(out[k], in[i][j])
		}
		k++
	}

	return out
}

func CopyMap[T comparable, U any](in map[T]U) map[T]U {
	out := make(map[T]U)
	for k := range in {
		out[k] = in[k]
	}

	return out
}

func StrToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return r
}

func Gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	return Gcd(b%a, a)
}

func PyMod(a int, b int) int {
	return (a%b + b) % b
}

func GetPermutations[T comparable](arr []T) [][]T {
	var heaps func(a []T, n int)
	res := [][]T{}

	heaps = func(a []T, n int) {
		if n == 1 {
			tmp := []T{}
			tmp = append(tmp, a...)
			res = append(res, tmp)
		} else {
			for i := range n {
				heaps(a, n-1)
				if n%2 != 0 {
					a[0], a[n-1] = a[n-1], a[0]
				} else {
					a[i], a[n-1] = a[n-1], a[i]
				}
			}
		}
	}
	heaps(arr, len(arr))
	return res
}
