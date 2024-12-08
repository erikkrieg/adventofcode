package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day7 struct {
	data []string
}

func (d *Day7) Setup() {
	data := input.Lines("day-7")
	if useTestInput {
		data = input.Lines("test-7")
	}
	d.data = data
}

func (d *Day7) Solve() {
	fmt.Println("Day 7")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day7) Part1() int {
	return 0
}

func (d *Day7) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[6] = (&Day7{}).Solve
}
