package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day01() {
	fmt.Println("- Day 01")
	file, err := os.Open("input/day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	part1(scanner)
}

func part1(scanner *bufio.Scanner) {
	fmt.Print("  - Part 1: ")
	currentCalorieSum := 0
	largestCalorieSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalorieSum > largestCalorieSum {
				largestCalorieSum = currentCalorieSum
			}
			currentCalorieSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCalorieSum += calories
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Most calories is ", largestCalorieSum, "\n")
}
