package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day6 struct {
	grid *lib.Grid[rune]
}

func (d *Day6) Setup() {
	grid := lib.NewGrid(input.LinesChars("day-6"))
	if useTestInput {
		grid = lib.NewGrid(input.LinesChars("test-6"))
	}
	d.grid = grid
}

func (d *Day6) Solve() {
	fmt.Println("Day 6")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day6) Part1() int {
	guard := d.grid.Find('^')
	guardDirIndex := 0
	visited := make(map[string]bool)
	for d.grid.Contains(guard) {
		pointId := fmt.Sprintf("%d,%d", guard.X, guard.Y)
		visited[pointId] = true
		nextGuard := d.grid.Relative(guard, lib.Directions[guardDirIndex])
		if nextGuard == nil {
			break
		} else if d.grid.Value(nextGuard) == '#' {
			guardDirIndex = (guardDirIndex + 1) % 4
			continue
		}
		guard = nextGuard
	}
	return len(visited)
}

func (d *Day6) Part2() int {
	// keep track of each visited point where the guard turned
	// each time the guard is about to take a step (next step), check if placing
	// a wall at that location would lead the guard to a previously visited turn point
	return 0
}

func init() {
	puzzleSolutions[5] = (&Day6{}).Solve
}
