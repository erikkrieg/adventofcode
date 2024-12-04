package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day4 struct {
	data [][]rune
	grid *lib.Grid[rune]
}

func (d *Day4) Setup() {
	data := input.LinesChars("day-4")
	if useTestInput {
		data = input.LinesChars("test-4")
	}
	d.data = data
	d.grid = lib.NewGrid(data)
}

func (d *Day4) Solve() {
	fmt.Println("Day 4")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day4) Part1() int {
	directions := []lib.Point{
		lib.Point{X: -1, Y: 0},  // Left
		lib.Point{X: 1, Y: 0},   // Right
		lib.Point{X: 0, Y: -1},  // Up
		lib.Point{X: 0, Y: 1},   // Down
		lib.Point{X: -1, Y: -1}, // Up & Left
		lib.Point{X: 1, Y: -1},  // Up & Right
		lib.Point{X: -1, Y: 1},  // Down & Left
		lib.Point{X: 1, Y: 1},   // Down & Right
	}
	xmasOccurrences := 0
	for _, x := range d.grid.FindAll('X') {
		for _, dir := range directions {
			if d.checkLine(x, &dir) {
				xmasOccurrences += 1
			}
		}
	}
	return xmasOccurrences
}

func (d *Day4) checkLine(current *lib.Point, delta *lib.Point) bool {
	currentValue := d.grid.Value(current)
	if currentValue == 'S' {
		return true
	}
	var target rune
	switch currentValue {
	case 'X':
		target = 'M'
	case 'M':
		target = 'A'
	case 'A':
		target = 'S'
	}
	next := d.grid.Relative(current, delta)
	if next == nil {
		return false
	}
	if d.grid.Value(next) == target {
		return d.checkLine(next, delta)
	}
	return false
}

func (d *Day4) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[3] = (&Day4{}).Solve
}
