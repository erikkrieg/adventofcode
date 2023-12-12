package solutions

import (
	"fmt"
	"slices"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[10] = day11Solution
}

func setupDay11() [][]rune {
	data := input.LinesChars("day-11")
	if useTestInput {
		data = input.LinesChars("test-11")
	}
	return data
}

func NewExpandedSpace(data [][]rune) *lib.Grid[rune] {
	expanded := [][]rune{}
	for _, row := range data {
		expanded = append(expanded, row)
		if !slices.Contains(row, '#') {
			expanded = append(expanded, slices.Clone(row))
		}
	}

	for x := 0; x < len(expanded[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(expanded); y++ {
			hasGalaxy = expanded[y][x] == '#'
			if hasGalaxy {
				break
			}
		}
		if hasGalaxy {
			continue
		}
		for y := 0; y < len(expanded); y++ {
			expanded[y] = slices.Insert(expanded[y], x, '.')
		}
		x++
	}

	return lib.NewGrid(expanded)
}

func DebugSpace(grid *lib.Grid[rune]) {
	for _, r := range grid.Points {
		for _, s := range r {
			fmt.Printf("%s ", string(s))
		}
		fmt.Printf("\n")
	}
}

func day11Solution() {
	fmt.Println("Day 11")
	data := setupDay11()
	grid := NewExpandedSpace(data)
	DebugSpace(grid)
	galaxies := grid.FindAll('#')
	sum := 0
	for _, a := range galaxies {
		for _, b := range galaxies {
			sum += a.Distance(b)
		}
	}
	Solution{
		Part1: sum / 2,
		Part2: nil,
	}.Print()
}
