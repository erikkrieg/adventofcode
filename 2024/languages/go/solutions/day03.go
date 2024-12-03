package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day3 struct {
	data []string
}

func (d *Day3) Setup() {
	data := input.Lines("day-3")
	if useTestInput {
		data = input.Lines("test-3")
	}
	d.data = data
}

func (d *Day3) Solve() {
	fmt.Println("Day 3")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day3) Part1() int {
	return 0
}

func (d *Day3) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[2] = (&Day3{}).Solve
}
