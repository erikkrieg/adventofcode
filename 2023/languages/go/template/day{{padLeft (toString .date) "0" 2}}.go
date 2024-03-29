package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[{{sub (atoi .date) 1}}] = day{{.date}}Solution
}

func setupDay{{.date}}() []string {
	data := input.Lines("day-{{.date}}")
	if useTestInput {
		data = input.Lines("test-{{.date}}")
	}
	return data
}

func day{{.date}}Solution() {
	fmt.Println("Day {{.date}}")
	data := setupDay{{.date}}()
	Solution{
		Part1: nil,
		Part2: nil,
	}.Print()
}
