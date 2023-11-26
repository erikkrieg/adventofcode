package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/solutions"
)

func main() {
	fmt.Println("AOC 2023")

	puzzleDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	solutions.Solve(puzzleDay)

	// Examples of importing inputs as varied data structures.
	chars := input.Chars("day-1")
	fmt.Printf("Chars: %+v\n", chars)

	lines := input.Lines("day-1")
	fmt.Printf("Lines: %+v\n", lines)

	linesChars := input.LinesChars("day-1")
	fmt.Printf("Chars per line: %+v\n", linesChars)
}
