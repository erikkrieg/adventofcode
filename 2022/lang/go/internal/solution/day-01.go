package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	file, err = os.Open("input/day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)
	part2(scanner)
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

func part2(scanner *bufio.Scanner) {
	fmt.Print("  - Part 2: ")
	calorieSums := []int{}
	currentCalorieSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calorieSums = append(calorieSums, currentCalorieSum)
			currentCalorieSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCalorieSum += calories
		}
	}
	sort.Ints(calorieSums)
	elfCount := len(calorieSums)
	topThree := calorieSums[elfCount-1] + calorieSums[elfCount-2] + calorieSums[elfCount-3]
	fmt.Print("Sum of calories for top 3 elves is ", topThree, "\n")
}
