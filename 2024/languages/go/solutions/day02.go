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
	for _, report := range d.data {
		levels := strings.Fields(report)
		if d.findUnsafeLevel(levels) >= 0 {
			continue
		}
		totalSafeReports += 1
	}
	return totalSafeReports
}

func (d *Day2) Part2() int {
	totalSafeReports := 0
	for _, report := range d.data {
		levels := strings.Fields(report)
		firstUnsafeLevelIndex := d.findUnsafeLevel(levels)
		if firstUnsafeLevelIndex >= 0 {
			foundSafeReport := false
			for i := range levels {
				l, _ := removeIndex(levels, i)
				foundSafeReport = d.findUnsafeLevel(l) == -1
				if foundSafeReport {
					break
				}
			}
			if !foundSafeReport {
				continue
			}
		}
		totalSafeReports += 1
	}
	return totalSafeReports
}

func removeIndex[T any](slice []T, i int) ([]T, error) {
	if i >= len(slice) || i < 0 {
		return nil, fmt.Errorf("Index is out of bounds: %d", i)
	}
	s := make([]T, 0, len(slice))
	s = append(s, slice[:i]...)
	s = append(s, slice[i+1:]...)
	return s, nil
}

func (d *Day2) findUnsafeLevel(levels []string) int {
	var previousLevel, direction *int
	for i, level := range levels {
		level := lib.Atoi(level)
		if previousLevel != nil {
			diff := (*previousLevel - level)
			diffDir := lib.Sign(diff)
			if direction == nil {
				direction = &diffDir
			}
			unsafeDiff := (lib.Abs(diff) < 1 || lib.Abs(diff) > 3)
			unsafeDir := *direction != diffDir
			if unsafeDiff || unsafeDir {
				return i - 1
			}

		}
		previousLevel = &level
	}
	return -1
}

func init() {
	puzzleSolutions[1] = (&Day2{}).Solve
}
