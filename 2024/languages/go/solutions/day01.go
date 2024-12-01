package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
)

func init() {
	puzzleSolutions[0] = day1Solution
}

func setupDay1() []string {
	data := input.Lines("day-1")
	if useTestInput {
		data = input.Lines("test-1")
	}
	return data
}

func day1Solution() {
	fmt.Println("Day 1")
	data := setupDay1()
	fmt.Printf("Data: %v\n", data)
	Solution{
		Part1: nil,
		Part2: nil,
	}.Print()
}
