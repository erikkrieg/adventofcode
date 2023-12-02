package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/erikkrieg/adventofcode/2023/solutions"
)

func main() {
	fmt.Println("AOC 2023")

	puzzleDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	solutions.Solve(puzzleDay, os.Getenv("TEST") == "true")
}
