package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[7] = day8Solution
}

func setupDay8() []string {
	data := input.Lines("day-8")
	if useTestInput {
		data = input.Lines("test-8")
	}
	return data
}

func day8Solution() {
	fmt.Println("Day 8")
	data := setupDay8()
	network := parseNetwork(data)

	Solution{
		Part1: day8Part1(data[0], network),
		Part2: nil,
	}.Print()
}

func day8Part1(moves string, network Network) int {
	node := "AAA"
	count := 0
	size := len(moves)
	for {
		move := moves[count%size]
		if node == "ZZZ" {
			break
		}
		next := 0
		if move == 'R' {
			next = 1
		}
		node = network[node][next]
		count++
	}
	return count
}

type Network map[string][2]string

func parseNetwork(data []string) Network {
	network := make(map[string][2]string)
	for _, d := range data[2:] {
		key := d[:3]
		nodes := [2]string{
			d[7:10],
			d[12:15],
		}
		network[key] = nodes
	}
	return network
}
