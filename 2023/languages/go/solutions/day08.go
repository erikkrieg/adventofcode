package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
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

func day8Solution() {
	fmt.Println("Day 8")
	part1Data := setupDay8("1")
	part2Data := setupDay8("2")

	if try(part2Data[0], parseNetwork(part2Data)) == 1 {
		return
	}

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

func try(moves string, network Network) int {
	// do these get into loops?
	// if they get into loops, how often do they land on Z
	// Is there a least common factor dynamic in this problem?
	//
	// So, if there is some eventual loop, then the time it takes the last thread to reach it might matter?
	tailNodes := []string{}
	visited := []map[string]bool{}
	for node := range network {
		if node[2] == 'A' {
			tailNodes = append(tailNodes, node)
			visited = append(visited, make(map[string]bool))
		}
	}

	fmt.Printf("%+v\n", tailNodes)

	size := len(moves)
	count := 0
	loopIntervals := make([]int, len(tailNodes))
	skips := 0
	for {
		move := moves[count%size]
		count++
		if skips > len(tailNodes) {
			break
		}
		for i, node := range tailNodes {
			if loopIntervals[i] > 0 {
				skips++
				continue
			}
			skips = 0
			key := fmt.Sprintf("%s-%d", node, count%size)
			if visited[i][key] {
				fmt.Printf("%d loop started at move: %d, pos: %s, key: %s\n", i, count, node, key)
				loopIntervals[i] = count
			} else {
				visited[i][key] = true
			}
			next := 0
			if move == 'R' {
				next = 1
			}
			tailNodes[i] = network[node][next]
		}
	}

	return 1
}

func day8Part2(moves string, network Network) int {
	tailNodes := []string{}
	for node := range network {
		if node[2] == 'A' {
			tailNodes = append(tailNodes, node)
		}
	}
	size := len(moves)
	count := 0
	for {
		move := moves[count%size]
		for i, node := range tailNodes {
			next := 0
			if move == 'R' {
				next = 1
			}
			tailNodes[i] = network[node][next]
		}
		count++
		endInZ := true
		for _, n := range tailNodes {
			if n[2] != 'Z' {
				endInZ = false
				break
			}
		}
		if endInZ {
			break
		}
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
