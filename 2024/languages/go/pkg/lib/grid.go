package lib

type Grid[T comparable] struct {
	Points [][]T
	MaxY   int
	MaxX   int
}

func NewGrid[T comparable](points [][]T) *Grid[T] {
	maxY := len(points)
	maxX := len(points[0])
	return &Grid[T]{Points: points, MaxY: maxY, MaxX: maxX}
}

func (g *Grid[comparable]) Contains(point *Point) bool {
	return point.X >= 0 && point.X < g.MaxX && point.Y >= 0 && point.Y < g.MaxY
}

func (g *Grid[comparable]) Value(point *Point) comparable {
	var value comparable
	if g.Contains(point) {
		value = g.Points[point.Y][point.X]
	}
	return value
}

func (g *Grid[comparable]) find(value comparable, all bool) []*Point {
	var matches []*Point
	for y, row := range g.Points {
		for x, v := range row {
			if v == value {
				matches = append(matches, &Point{Y: y, X: x})
				if !all {
					return matches
				}
			}
		}
	}
	return matches
}

func (g *Grid[comparable]) FindAll(value comparable) []*Point {
	return g.find(value, true)
}

func (g *Grid[comparable]) Find(value comparable) *Point {
	matches := g.find(value, false)
	if len(matches) == 0 {
		return nil
	}
	return matches[0]
}

func (g *Grid[comparable]) relative(point *Point, dir *Point) *Point {
	p := &Point{X: point.X + dir.X, Y: point.Y + dir.Y}
	if !g.Contains(p) {
		return nil
	}
	return p
}

func (g *Grid[comparable]) Relative(point *Point, dir *Point) *Point {
	return g.relative(point, dir)
}

func (g *Grid[comparable]) Above(point *Point) *Point {
	return g.relative(point, &Point{X: 0, Y: -1})
}
func (g *Grid[comparable]) Below(point *Point) *Point {
	return g.relative(point, &Point{X: 0, Y: 1})
}

func (g *Grid[comparable]) Left(point *Point) *Point {
	return g.relative(point, &Point{X: -1, Y: 0})
}

func (g *Grid[comparable]) LeftAbove(point *Point) *Point {
	return g.relative(point, &Point{X: -1, Y: -1})
}

func (g *Grid[comparable]) LeftBelow(point *Point) *Point {
	return g.relative(point, &Point{X: -1, Y: -1})
}

func (g *Grid[comparable]) Right(point *Point) *Point {
	return g.relative(point, &Point{X: 1, Y: 0})
}

func (g *Grid[comparable]) RightAbove(point *Point) *Point {
	return g.relative(point, &Point{X: 1, Y: -1})
}

func (g *Grid[comparable]) RightBelow(point *Point) *Point {
	return g.relative(point, &Point{X: 1, Y: 1})
}

func (g *Grid[comparable]) Around(point *Point) []*Point {
	points := []*Point{}
	directions := [][]int{
		{-1, 0, 1},
		{-1, 0, 1},
	}
	for _, x := range directions[0] {
		for _, y := range directions[1] {
			if x == 0 && y == 0 {
				continue
			}
			p := g.relative(point, &Point{X: x, Y: y})
			if p != nil {
				points = append(points, p)
			}
		}
	}
	return points
}
