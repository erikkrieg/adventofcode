package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/erikkrieg/adventofcode-2022/internal/solution"
)

func main() {
	fmt.Println("Advent of Code 2022")
	puzzleDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	switch puzzleDay {
	case 1:
		solution.Day01()
	default:
		log.Fatal("Solution not found: ", puzzleDay)
	}
}
