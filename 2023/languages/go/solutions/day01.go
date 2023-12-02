package solutions

import (
	"fmt"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

var day1Part1Test = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var day1Part2Test = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func init() {
	puzzleSolutions[0] = day1Solution
}

func day1Solution() {
	fmt.Println("Day 1")
	digits := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	fmt.Printf(" Part 1: %d\n", day1BothParts(digits[:9], &day1Part1Test))
	fmt.Printf(" Part 2: %d\n", day1BothParts(digits, &day1Part2Test))
}

func day1BothParts(digits []string, test *[]string) int {
	documents := input.Lines("day-1")
	if useTestInput {
		documents = *test
	}
	return tre(documents, digits)
}

func tre(documents []string, digits []string) int {
	sum := 0
	for _, doc := range documents {
		sum += findFirstDigit(doc, digits)*10 + findLastDigit(doc, digits)
	}
	return sum
}

func findFirstDigit(doc string, digits []string) int {
	for i := 0; i < len(doc); i++ {
		for v, d := range digits {
			if strings.Contains(doc[:i+1], d) {
				return v%9 + 1
			}
		}
	}
	return 0
}
func findLastDigit(doc string, digits []string) int {
	for i := len(doc) - 1; i >= 0; i-- {
		for v, d := range digits {
			if strings.Contains(doc[i:], d) {
				return v%9 + 1
			}
		}
	}
	return 0
}
