package solutions

import (
	"fmt"
	"regexp"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
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
	productSum := 0
	for _, program := range d.data {
		r, _ := regexp.Compile(`mul\((\d*),(\d*)\)`)
		matches := r.FindAllStringSubmatch(program, -1)
		for _, m := range matches {
			productSum += lib.Atoi(m[1]) * lib.Atoi(m[2])
		}
	}
	return productSum
}

func (d *Day3) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[2] = (&Day3{}).Solve
}
