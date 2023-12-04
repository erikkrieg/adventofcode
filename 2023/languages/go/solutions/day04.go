package solutions

import (
	"fmt"
	"math"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

var day4Test = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func init() {
	puzzleSolutions[3] = day4Solution
}

func day4Solution() {
	fmt.Println("Day 4")
	day4BothParts().Print()
}

func day4BothParts() Solution {
	cards := input.Lines("day-4")
	if useTestInput {
		cards = day4Test
	}
	cardCount := make([]int, len(cards))
	points := 0
	for i, card := range cards {
		cardCount[i] += 1
		numbers := strings.Split(strings.Split(card, ":")[1], "|")
		winningNums := make(map[string]bool)
		for _, n := range strings.Split(numbers[0], " ") {
			if n != "" {
				winningNums[n] = true
			}
		}
		ourNums := strings.Split(numbers[1], " ")
		matches := 0
		for _, n := range ourNums {
			if winningNums[n] {
				matches++
			}
		}
		for y := i + 1; y < len(cards) && y <= i+matches; y++ {
			cardCount[y] += cardCount[i]
		}
		points += int(math.Pow(2.0, float64(matches-1)))
	}
	cardTotal := 0
	for _, c := range cardCount {
		cardTotal += c
	}
	return Solution{
		Part1: points,
		Part2: cardTotal,
	}
}
