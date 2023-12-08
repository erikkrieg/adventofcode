package solutions

import "fmt"

func init() {
	puzzleSolutions[{{sub (atoi .date) 1}}] = day{{.date}}Solution
}

func setupDay{{.date}}() []string {
	data := input.Lines("day-7")
	if useTestInput {
		data = input.Lines("test-7")
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
