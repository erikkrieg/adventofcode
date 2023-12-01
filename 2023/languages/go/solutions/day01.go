package solutions

import (
	"fmt"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[0] = day1Solution
}

func day1Solution() {
	fmt.Println("Day 1")
	day1Part1(false)
	day1Part2(false)
}

func day1Part1(test bool) {
	fmt.Print(" Part 1: ")
	documents := input.Lines("day-1")
	if test {
		documents = []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}
	}
	digits := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	calibartion := trebuchet(documents, digits)
	fmt.Println(calibartion)
}

func day1Part2(test bool) {
	fmt.Print(" Part 2: ")
	documents := input.Lines("day-1")
	if test {
		documents = []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}
	}
	digits := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	calibartion := trebuchet(documents, digits)
	fmt.Println(calibartion)
}

func trebuchet(documents []string, digits []string) int {
	sum := 0
	for _, doc := range documents {
		first := len(doc)
		last := -1
		nums := [2]int{}
		for _, d := range digits {
			i := strings.Index(doc, d)
			if i > -1 && i < first {
				first = i
				nums[0] = toInt(d)
			}
			i = strings.LastIndex(doc, d)
			if i > last {
				last = i
				nums[1] = toInt(d)
			}
		}
		sum += nums[0]*10 + nums[1]
	}
	return sum
}

func toInt(slice string) int {
	var num int
	switch slice {
	case "one", "1":
		num = 1
	case "two", "2":
		num = 2
	case "three", "3":
		num = 3
	case "four", "4":
		num = 4
	case "five", "5":
		num = 5
	case "six", "6":
		num = 6
	case "seven", "7":
		num = 7
	case "eight", "8":
		num = 8
	case "nine", "9":
		num = 9
	}
	return num
}
