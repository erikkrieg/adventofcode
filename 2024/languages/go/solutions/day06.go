package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day6 struct {
	data []string
}

func (d *Day6) Setup() {
	data := input.Lines("day-6")
	if useTestInput {
		data = input.Lines("test-6")
	}
	d.data = data
}

func (d *Day6) Solve() {
	fmt.Println("Day 6")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day6) Part1() int {
	return 0
}

func (d *Day6) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[5] = (&Day6{}).Solve
}
