package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[10] = day11Solution
}

func setupDay11() []string {
	data := input.Lines("day-11")
	if useTestInput {
		data = input.Lines("test-11")
	}
	return data
}

func day11Solution() {
	fmt.Println("Day 11")
	data := setupDay11()
	Solution{
		Part1: nil,
		Part2: nil,
	}.Print()
}
