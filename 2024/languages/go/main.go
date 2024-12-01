package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/erikkrieg/adventofcode/2024/solutions"
)

func main() {
	fmt.Println("AOC 2024")

	puzzleDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	solutions.Solve(puzzleDay, os.Getenv("TEST") == "true")
}
