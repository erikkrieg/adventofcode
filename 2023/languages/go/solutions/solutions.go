package solutions

import (
	"log"
)

type PuzzleFunc func()

var puzzleSolutions [25]PuzzleFunc

func Solve(day int) {
	solution := puzzleSolutions[day-1]
	if solution == nil {
		log.Fatalf("No solution found for day %d", day)
	}
	solution()
}
