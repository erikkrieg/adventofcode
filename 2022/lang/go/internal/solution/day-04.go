package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day04() {
	fmt.Println("- Day 04")

	file, err := os.Open("input/day-04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fullOverlapCount := 0
	partOverlapCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := [][]int{}
		for i, s := range strings.Split(line, ",") {
			assignments = append(assignments, []int{})
			for _, ss := range strings.Split(s, "-") {
				si, err := strconv.Atoi(ss)
				if err != nil {
					log.Fatal(err)
				}
				assignments[i] = append(assignments[i], si)
			}
		}

		a := assignments[0]
		b := assignments[1]
		if rangeInRange(a, b) || rangeInRange(b, a) {
			fullOverlapCount += 1
			partOverlapCount += 1
		} else if inRange(a[0], b) || inRange(a[1], b) || inRange(b[0], a) || inRange(b[1], a) {
			partOverlapCount += 1
		}
	}

	fmt.Println("  - Part 1: ", fullOverlapCount)
	fmt.Println("  - Part 2: ", partOverlapCount)
}

func inRange(n int, r []int) bool {
	return n >= r[0] && n <= r[1]
}

func rangeInRange(a, b []int) bool {
	return inRange(a[0], b) && inRange(a[1], b)
}
