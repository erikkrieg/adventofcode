package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

var day3Test = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"..........",
	".........+",
	".........1",
	"..........",
	"..........",
	"......755.",
	"...$.*....",
	".664.598..",
}

func init() {
	puzzleSolutions[2] = day3Solution
}

func day3Solution() {
	fmt.Println("Day 3")
	fmt.Printf(" - Part 1: %d\n", day3Part1(day3Test))
}

func day3Part1(schematic []string) int {
	digits := "0123456789"
	if !useTestInput {
		schematic = input.Lines("day-3")
	}
	partNumSum := 0
	for y, s := range schematic {
		num := ""
		for x, char := range s {
			isDigit := strings.ContainsRune(digits, char)
			if isDigit {
				num += string(char)
			}
			if (!isDigit || x == len(s)-1) && num != "" {
				isPart := false
				xmin := x - len(num) - 1
				if xmin < 0 {
					xmin = 0
				}
				xmax := x + 1
				if xmax > len(s) {
					xmax -= 1
				}
				around := []string{}
				if y > 0 {
					around = append(around, schematic[y-1][xmin:xmax])
				}
				if y < len(schematic)-1 {
					around = append(around, schematic[y+1][xmin:xmax])
				}
				if xmin-1 > 0 {
					around = append(around, s[xmin:xmin+1])
				}
				if xmax <= len(s) {
					around = append(around, s[x:xmax])
				}
				fmt.Printf("xmin: %d, xmax: %d\n", xmin, xmax)
				fmt.Printf("schema: %s -> around: %d %+v\n", num, len(around), around)
				for _, side := range around {
					if strings.ContainsAny(side, "!@#$%^&*-=+/:") {
						isPart = true
						break
					}
				}
				if isPart {
					fmt.Printf("Num: %s is a part\n\n", num)
					n, err := strconv.Atoi(num)
					if err != nil {
						panic(err)
					}
					partNumSum += n
				} else {
					fmt.Println()
				}
				num = ""
			}

		}
	}
	return partNumSum
}
