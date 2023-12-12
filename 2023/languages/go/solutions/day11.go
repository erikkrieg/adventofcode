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
	Solution{
		Part1: day11Part1(setupDay11()),
		Part2: day11Part2(setupDay11()),
	}.Print()
}

func day11Part1(data [][]rune) int {
	grid := NewExpandedSpace(data)
	galaxies := grid.FindAll('#')
	sum := 0
	for _, a := range galaxies {
		for _, b := range galaxies {
			sum += a.Distance(b)
		}
	}
	return sum / 2
}

func day11Part2(data [][]rune) int {
	grid := lib.NewGrid(data)
	emptyY, emptyX := emptySpaces(data)
	galaxies := grid.FindAll('#')
	sum := 0
	for _, a := range galaxies {
		for _, b := range galaxies {
			sum += a.Distance(b)
			sum += fillBetween(a.X, b.X, emptyX) + fillBetween(a.Y, b.Y, emptyY)
		}
	}
	return sum / 2
}

func fillBetween(a, b int, empty map[int]bool) int {
	mul := 1_000_000
	count := 0
	for k := range empty {
		if min(a, b) < k && max(a, b) > k {
			count++
		}
	}
	return (mul - 1) * count
}

func emptySpaces(data [][]rune) (rows, cols map[int]bool) {
	rows = make(map[int]bool)
	cols = make(map[int]bool)
	for y, row := range data {
		if !slices.Contains(row, '#') {
			rows[y] = true
		}
	}
	for x := 0; x < len(data[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(data); y++ {
			hasGalaxy = data[y][x] == '#'
			if hasGalaxy {
				break
			}
		}
		if !hasGalaxy {
			cols[x] = true
		}
	}
	return
}
