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
	wellStockedElves := elvesWithMostCalories(scanner, 3)
	part1(wellStockedElves)
	part2(wellStockedElves)
}

func elvesWithMostCalories(scanner *bufio.Scanner, elfLimit int) []int {
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
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))
	return calorieSums[:elfLimit]
}

func part1(topElves []int) {
	fmt.Println("  - Part 1: Most calories is ", topElves[0])
}

func part2(topElves []int) {
	sum := topElves[0] + topElves[1] + topElves[2]
	fmt.Println("  - Part 2: Sum of calories for top 3 elves is ", sum)
}
