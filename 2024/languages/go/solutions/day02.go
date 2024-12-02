package solutions

import (
	"fmt"
	"strings"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day2 struct {
	data []string
}

func (d *Day2) Setup() {
	data := input.Lines("day-2")
	if useTestInput {
		data = input.Lines("test-2")
	}
	d.data = data
}

func (d *Day2) Solve() {
	fmt.Println("Day 2")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day2) Part1() int {
	totalSafeReports := 0
reportLoop:
	for _, report := range d.data {
		levelDifference := 0
		levels := strings.Fields(report)
		var previousLevel *int
		for _, level := range levels {
			level := lib.Atoi(level)
			if previousLevel != nil {
				diff := (*previousLevel - level)
				nextLevelDifference := levelDifference + diff
				if lib.Abs(diff) < 1 || lib.Abs(diff) > 3 {
					// This level is not safe because the distance between two levels
					// is outside of the designated bounds (1-3 inclusive).
					continue reportLoop
				}
				if lib.Abs(nextLevelDifference) <= lib.Abs(levelDifference) {
					// This level is not safe because it is not always increasing or decreasing
					continue reportLoop
				}
				levelDifference = nextLevelDifference
			}
			previousLevel = &level
		}
		totalSafeReports += 1
	}
	return totalSafeReports
}

func (d *Day2) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[1] = (&Day2{}).Solve
}
