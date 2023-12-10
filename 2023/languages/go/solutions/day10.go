package solutions

import (
	"fmt"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[9] = day10Solution
}

func setupDay10() [][]rune {
	data := input.LinesChars("day-10")
	if useTestInput {
		data = input.LinesChars("test-10.1")
	}
	return data
}

func day10Solution() {
	fmt.Println("Day 10")
	data := setupDay10()
	Solution{
		Part1: day10Part1(data),
		Part2: nil,
	}.Print()
}

type Pipe struct {
	grid  *lib.Grid[rune]
	point *lib.Point
	kind  rune
}

func NewPipe(grid *lib.Grid[rune], point *lib.Point) *Pipe {
	return &Pipe{grid: grid, point: point, kind: grid.Value(point)}
}

func (p *Pipe) Connections() []*Pipe {
	connections := []*Pipe{}
	switch p.kind {
	case '|':
		connections = append(connections, NewPipe(p.grid, p.grid.Above(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Below(p.point)))
	case '-':
		connections = append(connections, NewPipe(p.grid, p.grid.Left(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Right(p.point)))
	case 'L':
		connections = append(connections, NewPipe(p.grid, p.grid.Above(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Right(p.point)))
	case 'J':
		connections = append(connections, NewPipe(p.grid, p.grid.Above(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Left(p.point)))
	case '7':
		connections = append(connections, NewPipe(p.grid, p.grid.Left(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Below(p.point)))
	case 'F':
		connections = append(connections, NewPipe(p.grid, p.grid.Right(p.point)))
		connections = append(connections, NewPipe(p.grid, p.grid.Below(p.point)))
	case 'S':
		for _, ap := range p.grid.Around(p.point) {
			pipe := NewPipe(p.grid, ap)
			for _, c := range pipe.Connections() {
				if c.kind == 'S' {
					connections = append(connections, pipe)
				}
			}
		}
	case '.': // Ignore "ground"
	}
	return connections
}

func (p *Pipe) move(moveCount int, visited *map[lib.Point]int) {
	mc, ok := (*visited)[*p.point]
	if !ok || mc > moveCount {
		(*visited)[*p.point] = moveCount
		for _, next := range p.Connections() {
			next.move(moveCount+1, visited)
		}
	}
}

func day10Part1(data [][]rune) int {
	grid := lib.NewGrid(data)
	start := Pipe{point: grid.Find('S'), kind: 'S', grid: grid}
	visited := make(map[lib.Point]int)
	start.move(0, &visited)
	futhestPipe := 0
	for _, v := range visited {
		futhestPipe = max(futhestPipe, v)
	}
	return futhestPipe
}
