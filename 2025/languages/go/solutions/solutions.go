package solutions

import (
	"fmt"
	"log"
)

type Solution struct {
	Part1 interface{}
	Part2 interface{}
}

func (s Solution) Print() {
	fmt.Printf("  - part 1: %v\n  - part 2: %v\n", s.Part1, s.Part2)
}

type PuzzleFunc func()

var puzzleSolutions [25]PuzzleFunc
var useTestInput bool

func Solve(day int, test bool) {
	useTestInput = test
	solution := puzzleSolutions[day-1]
	if solution == nil {
		log.Fatalf("No solution found for day %d", day)
	}
	solution()
}
