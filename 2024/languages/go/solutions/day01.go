package solutions

import (
	"fmt"
	"sort"
	"strings"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day1 struct {
	data []string
}

func (d *Day1) Setup() {
	data := input.Lines("day-1")
	if useTestInput {
		data = input.Lines("test-1")
	}
	d.data = data
}

func (d *Day1) Solve() {
	fmt.Println("Day 1")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day1) Part1() int {
	var distanceSum int
	left := make([]int, len(d.data))
	right := make([]int, len(d.data))
	for i, line := range d.data {
		fields := strings.Fields(line)
		left[i] = lib.Atoi(fields[0])
		right[i] = lib.Atoi(fields[1])
	}
	sort.Ints(left)
	sort.Ints(right)
	for i := 0; i < len(d.data); i++ {
		distanceSum += lib.Abs(left[i] - right[i])
	}
	return distanceSum
}

func (d *Day1) Part2() int {
	left := make([]int, len(d.data))
	right := make(map[int]int)
	for i, line := range d.data {
		fields := strings.Fields(line)
		left[i] = lib.Atoi(fields[0])
		right[lib.Atoi(fields[1])] += 1
	}
	similarity := 0
	for _, n := range left {
		similarity += n * right[n]
	}
	return similarity
}

func init() {
	puzzleSolutions[0] = (&Day1{}).Solve
}
