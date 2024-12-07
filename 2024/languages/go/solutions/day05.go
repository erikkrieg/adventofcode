package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day5 struct {
	data []string
}

func (d *Day5) Setup() {
	data := input.Lines("day-5")
	if useTestInput {
		data = input.Lines("test-5")
	}
	d.data = data
}

func (d *Day5) Solve() {
	fmt.Println("Day 5")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day5) Part1() int {
	return 0
}

func (d *Day5) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[4] = (&Day5{}).Solve
}
