package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[7] = day8Solution
}

func setupDay8(part string) []string {
	data := input.Lines("day-8")
	if useTestInput {
		data = input.Lines("test-8." + part)
	}
	return data
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

func day8Solution() {
	fmt.Println("Day 8")
	part1Data := setupDay8("1")
	part2Data := setupDay8("2")
	Solution{
		Part1: day8Part1(part1Data[0], parseNetwork(part1Data)),
		Part2: day8Part2(part2Data[0], parseNetwork(part2Data)),
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

// After much tinkering I realized that each "thread" of nodes:
// - loops in a repeating patterns
// - encounters only one terminal node (**Z)
// - the number of steps to reach the terminal node for the first time is the interval
//
// Knowing this, I stripped out a bunch of code meant for handling cases where
// thre are multiple possible **Z nodes per thread as well as the case where threads
// do not completely loop front to back, but get stuck in smaller loops (which was not the case).
func day8Part2(moves string, network Network) int {
	tailNodes := []string{}
	for node := range network {
		if node[2] == 'A' {
			tailNodes = append(tailNodes, node)
		}
	}
	size := len(moves)
	count := 0
	intervals := make(map[int]int)
	for len(intervals) < len(tailNodes) {
		move := moves[count%size]
		for i, node := range tailNodes {
			if _, ok := intervals[i]; ok {
				continue
			}
			if node[2] == 'Z' {
				intervals[i] = count
			}
			next := 0
			if move == 'R' {
				next = 1
			}
			tailNodes[i] = network[node][next]
		}
		count++
	}
	f := make([]int, len(intervals))
	for i, v := range intervals {
		f[i] = v
	}
	return lib.IntsLCM(f...)
}
