package solutions

import (
	"log"
)

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
