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
		data = input.LinesChars("test-10.2")
	}
	return data
}

func day10Solution() {
	fmt.Println("Day 10")
	data := setupDay10()
	Solution{
		Part1: day10Part1(data),
		Part2: day10Part2(data),
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

type Direction = lib.Point

var DirLabel = map[lib.Point]string{
	{X: 0, Y: 0}:  "None",
	{X: 1, Y: 0}:  "Right",
	{X: -1, Y: 0}: "Left",
	{X: 0, Y: 1}:  "Down",
	{X: 0, Y: -1}: "Up",
}

// TODO: figure out edge case for detecting direction from start
func (p *Pipe) trace(prev lib.Point, edges *map[lib.Point]Direction) {
	if _, ok := (*edges)[*p.point]; ok {
		return
	}
	// TODO: Figure out why this rotation is not actually working
	newDir := lib.Point{X: p.point.X - prev.X, Y: p.point.Y - prev.Y}
	fmt.Printf("%+v  -  ", DirLabel[newDir])
	(*edges)[*p.point] = newDir
	for _, next := range p.Connections() {
		next.trace(*p.point, edges)
		if p.kind == 'S' {
			break
		}
	}
}

func (p *Pipe) spread(edges *map[lib.Point]Direction, area *map[lib.Point]bool) {
	_, isVisited := (*area)[*p.point]
	_, isEdge := (*edges)[*p.point]
	if isVisited || isEdge {
		return
	}
	(*area)[*p.point] = true
	for _, point := range p.grid.Around(p.point) {
		next := NewPipe(p.grid, point)
		next.spread(edges, area)
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

/* var debugPoints = []lib.Point{
	{Y: 3, X: 3}, {Y: 3, X: 6},
	{Y: 4, X: 1}, {Y: 4, X: 4}, {Y: 4, X: 5}, {Y: 4, X: 8},
	{Y: 5, X: 2}, {Y: 5, X: 3}, {Y: 5, X: 6}, {Y: 5, X: 7},
} */

func day10Part2(data [][]rune) int {
	grid := lib.NewGrid(data)
	start := Pipe{point: grid.Find('S'), kind: 'S', grid: grid}
	loop := make(map[lib.Point]Direction)
	start.trace(*start.point, &loop)

	sideA := make(map[lib.Point]bool)
	sideB := make(map[lib.Point]bool)
	//fmt.Printf("\n")
	for point, dir := range loop {
		pipe := NewPipe(grid, &point)
		//if pipe.kind != '|' && pipe.kind != '-' {
		if pipe.kind == 'S' {
			continue
		}

		sides := pipe.sides(dir)

		// TODO: remove debug code
		//if point.X == 1 && point.Y == 4 {
		/* if slices.Contains(debugPoints, point) {
			fmt.Printf("Hit %+v -> dir %+v\n", point, dir)
			for _, s := range sides[0] {
				fmt.Printf("  side A: %+v = %s\n", s.point, string(s.kind))
			}
			for _, s := range sides[1] {
				fmt.Printf("  side B: %+v = %s\n", s.point, string(s.kind))
			}
		} */

		for _, s := range sides[0] {
			if grid.Contains(s.point) {
				s.spread(&loop, &sideA)
			}
		}
		for _, s := range sides[1] {
			if grid.Contains(s.point) {
				s.spread(&loop, &sideB)
			}
		}
	}

	fmt.Println("Side A:")
	for p := range sideA {
		fmt.Printf("  %+v > %s\n", p, string(grid.Value(&p)))
	}

	fmt.Println("Side B:")
	for p := range sideB {
		fmt.Printf("  %+v > %s\n", p, string(grid.Value(&p)))
	}

	/* fmt.Println("Debug points")
	for _, dp := range debugPoints {
		fmt.Printf("Point: %+v, DirLabel: %s, Dir:%+v\n", dp, DirLabel[loop[dp]], loop[dp])
	} */

	fmt.Printf("\nTotal size: %d\n", grid.MaxX*grid.MaxY)
	fmt.Printf("Edges len: %d\n", len(loop))
	fmt.Printf("A len: %d, B len: %d\n\n", len(sideA), len(sideB))

	return 0
}

// This approach requires I figure out how to track distinct sides of the loop
// and then figure out which side is the inner or outer side. But I think that
// in order to do this I need to rotate.
func (p *Pipe) sides(dir lib.Point) [2][]*Pipe {
	var sideA, sideB []*Pipe

	switch p.kind {
	case '|':
		sideA = []*Pipe{NewPipe(p.grid, &lib.Point{X: p.point.X + dir.Y, Y: p.point.Y + dir.X})}
		sideB = []*Pipe{NewPipe(p.grid, &lib.Point{X: p.point.X - dir.Y, Y: p.point.Y - dir.X})}
	case '-':
		sideA = []*Pipe{NewPipe(p.grid, &lib.Point{X: p.point.X - dir.Y, Y: p.point.Y - dir.X})}
		sideB = []*Pipe{NewPipe(p.grid, &lib.Point{X: p.point.X + dir.Y, Y: p.point.Y + dir.X})}
	case '7':
		switch DirLabel[dir] {
		case "Up":
			sideB = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X + 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y - 1}),
			}
		case "Right":
			sideA = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X + 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y - 1}),
			}
		}
	case 'J':
		switch DirLabel[dir] {
		case "Right":
			sideB = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y + 1}),
				NewPipe(p.grid, &lib.Point{X: p.point.X + 1, Y: p.point.Y}),
			}
		case "Up":
			sideA = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y + 1}),
				NewPipe(p.grid, &lib.Point{X: p.point.X + 1, Y: p.point.Y}),
			}
		}
	case 'F':
		switch DirLabel[dir] {
		case "Left":
			sideB = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X - 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y - 1}),
			}
		case "Up":
			sideA = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X - 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y - 1}),
			}
		}
	case 'L':
		switch DirLabel[dir] {
		case "Down":
			sideB = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X - 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y + 1}),
			}
		case "Left":
			sideA = []*Pipe{
				NewPipe(p.grid, &lib.Point{X: p.point.X - 1, Y: p.point.Y}),
				NewPipe(p.grid, &lib.Point{X: p.point.X, Y: p.point.Y + 1}),
			}
		}
	}

	// TODO: remove
	/* if p.point.X == 1 && p.point.Y == 4 {
		fmt.Printf("sides hit\n")
		fmt.Printf("dir: %+v\n", dir)
	} */

	/* if dir.Y == -1 { // going up
	} else if dir.Y == 1 { // going down
	} else if dir.X == -1 { // going left
	} else if dir.X == 1 { // going right
	} */

	return [2][]*Pipe{sideA, sideB}
}
