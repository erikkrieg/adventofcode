package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

type Day{{.date}} struct {
	data []string
}

func (d *Day{{.date}}) Setup() {
	data := input.Lines("day-{{.date}}")
	if useTestInput {
		data = input.Lines("test-{{.date}}")
	}
	d.data = data
}

func (d *Day{{.date}}) Solve() {
	fmt.Println("Day {{.date}}")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day{{.date}}) Part1() int {
	return 0
}

func (d *Day{{.date}}) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[{{sub (atoi .date) 1}}] = (&Day{{.date}}{}).Solve
}
