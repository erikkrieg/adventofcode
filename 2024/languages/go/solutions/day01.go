package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

func init() {
	puzzleSolutions[0] = day1Solution
}

func setupDay1() []string {
	data := input.Lines("day-1")
	if useTestInput {
		data = input.Lines("test-1")
	}
	return data
}

func day1Solution() {
	fmt.Println("Day 1")
	data := setupDay1()
	Solution{
		Part1: day1Part1(data),
		Part2: nil,
	}.Print()
}

func day1Part1(data []string) int {
	var distanceSum int
	var left, right []int
	for _, d := range data {
		fields := strings.Fields(d)
		left = append(left, tryInt(fields[0]))
		right = append(right, tryInt(fields[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	for i := 0; i < len(data); i++ {
		distanceSum += lib.Abs(left[i] - right[i])
	}
	return distanceSum
}

func tryInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
