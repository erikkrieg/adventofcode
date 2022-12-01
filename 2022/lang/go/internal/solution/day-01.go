package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day01() {
	file, err := os.Open("input/day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentCalorieSum := 0
	largestCalorieSum := 0

	scanner := bufio.NewScanner(file)
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
	fmt.Println("Most calories is: ", largestCalorieSum)

}
