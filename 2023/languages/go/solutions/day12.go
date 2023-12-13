package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[11] = day12Solution
}

func setupDay12() []string {
	data := input.Lines("day-12")
	if useTestInput {
		data = input.Lines("test-12")
	}
	return data
}

func day12Solution() {
	fmt.Println("Day 12")
	Solution{
		Part1: day12Part1(setupDay12()),
		Part2: nil,
	}.Print()
}

func parseSprings(data string) (string, []int) {
	s := strings.Split(data, " ")
	springs := s[0]
	nums := lib.Map(strings.Split(s[1], ","), func(v string) int {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return i
	})
	return springs, nums
}

func day12Part1(data []string) int {
	combos := 0
	for _, d := range data {
		springs, nums := parseSprings(d)
		combos += rec(springs, nums)
	}
	return combos
}

func rec(springs string, nums []int) int {
	if !strings.Contains(springs, "?") {
		out := 0
		if validSprings(springs, nums) {
			out = 1
		}
		return out
	}
	a := strings.Replace(springs, "?", ".", 1)
	b := strings.Replace(springs, "?", "#", 1)
	return rec(a, nums) + rec(b, nums)
}

func validSprings(str string, groups []int) bool {
	springs := strings.Fields(strings.ReplaceAll(str, ".", " "))
	if len(springs) != len(groups) {
		return false
	}
	for i, s := range springs {
		if len(s) != groups[i] {
			return false
		}
	}
	return true
}
