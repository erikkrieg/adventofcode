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
		pointId := guard.Id()
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
	guard := d.grid.Find('^')
	guardDirIndex := 0
	visited := make(map[string]bool)
	loops := make(map[string]bool)
	for d.grid.Contains(guard) {
		pointId := guard.Id()
		visited[pointId] = true
		nextGuard := d.grid.Relative(guard, lib.Directions[guardDirIndex])
		if nextGuard == nil {
			break
		}
		nextGuardId := nextGuard.Id()
		if d.grid.Value(nextGuard) == '#' {
			guardDirIndex = (guardDirIndex + 1) % 4
			continue
		} else if _, ok := visited[nextGuardId]; !ok {
			checkIndex := (guardDirIndex + 1) % 4
			check := &lib.Point{X: guard.X, Y: guard.Y}
			checkVisited := make(map[string]bool)
			for {
				nextCheck := d.grid.Relative(check, lib.Directions[checkIndex])
				if nextCheck == nil {
					break
				}
				nextCheckId := nextCheck.Id()
				checkIntersectId := fmt.Sprintf("%s:%s", nextCheckId, check.Id())
				if _, ok := checkVisited[checkIntersectId]; ok {
					loops[nextGuardId] = true
					break
				}
				if d.grid.Value(nextCheck) == '#' || nextGuardId == nextCheckId {
					checkIndex = (checkIndex + 1) % 4
					continue
				}
				checkVisited[checkIntersectId] = true
				check = nextCheck
			}
		}
		guard = nextGuard
	}
	return len(loops)
}

func init() {
	puzzleSolutions[5] = (&Day6{}).Solve
}
