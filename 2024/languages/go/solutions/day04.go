package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day4 struct {
	data []string
}

func (d *Day4) Setup() {
	data := input.Lines("day-4")
	if useTestInput {
		data = input.Lines("test-4")
	}
	d.data = data
}

func (d *Day4) Solve() {
	fmt.Println("Day 4")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day4) Part1() int {
	return 0
}

func (d *Day4) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[3] = (&Day4{}).Solve
}
