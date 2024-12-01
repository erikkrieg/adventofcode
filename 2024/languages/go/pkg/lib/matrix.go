package lib

type Point struct {
	X, Y int
}

func (p *Point) Neighbors() []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y},     // Left
		{X: p.X - 1, Y: p.Y - 1}, // Top Left
		{X: p.X - 1, Y: p.Y + 1}, // Bottom Left

		{X: p.X + 1, Y: p.Y},     // Right
		{X: p.X + 1, Y: p.Y + 1}, // Bottom Right
		{X: p.X + 1, Y: p.Y - 1}, // Top Right

		{X: p.X, Y: p.Y - 1}, // Top
		{X: p.X, Y: p.Y + 1}, // Bottom
	}
}

func (p *Point) InBounds(x, y int) bool {
	return p.X > -1 && p.Y > -1 && p.X < x && p.Y < y
}

func (p *Point) Distance(other *Point) int {
	edge := Edge{Start: *p, End: *other}
	return edge.Len()
}

// Edge methods are designed around vertical and horizontal edges. Diagonal
// points will not be able to use all Edge methods.
type Edge struct {
	Start Point
	End   Point
}

func (e Edge) Len() int {
	deltaX := Abs(e.Start.X - e.End.X)
	deltaY := Abs(e.Start.Y - e.End.Y)
	return deltaX + deltaY
}

func (e Edge) SliceFrom(list [][]interface{}) []interface{} {
	if e.Start.Y == e.End.Y {
		return list[e.Start.Y][e.Start.X:e.End.X]
	} else if e.Start.X != e.End.X {
		panic("Edge SliceFrom method does not support diagonal edges.")
	}
	slice := make([]interface{}, e.Len())
	// This is currently assuming the end is later in the matrix. We can make this
	// more flexible later.
	i := 0
	for y := e.Start.Y; y <= e.End.Y; y++ {
		slice[i] = list[y][e.Start.X]
		i++
	}
	return slice
}
