package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day8 struct {
	data []string
}

func (d *Day8) Setup() {
	data := input.Lines("day-8")
	if useTestInput {
		data = input.Lines("test-8")
	}
	d.data = data
}

func (d *Day8) Solve() {
	fmt.Println("Day 8")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day8) Part1() int {
	return 0
}

func (d *Day8) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[7] = (&Day8{}).Solve
}
