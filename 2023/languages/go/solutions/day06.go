package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[5] = day6Solution
}

func day6Solution() {
	fmt.Println("Day 6")
	day6Part1().Print()
	day6Part2().Print()
}

type Race struct {
	Time     int
	Distance int
}

func day6Setup(oneRace bool) []Race {
	data := input.Lines("day-6")
	if useTestInput {
		data = input.Lines("test-6")
	}
	races := []Race{}
	for _, d := range data {
		split := strings.Split(d, ":")[1]
		values := []string{}
		if oneRace {
			values = append(values, strings.ReplaceAll(split, " ", ""))
		} else {
			values = strings.Fields(split)
		}
		for i, val := range values {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			if i < len(races) {
				races[i].Distance = v
			} else {
				races = append(races, Race{Time: v})
			}
		}
	}
	return races
}

func day6Part1() Solution {
	races := day6Setup(false)
	winningStrats := make([]int, len(races))
	for i, r := range races {
		for t := 1; t < r.Time; t++ {
			time := t * (r.Time - t)
			if time > r.Distance {
				winningStrats[i]++
			}
		}
	}
	part1 := 1
	for _, s := range winningStrats {
		part1 *= s
	}
	return Solution{
		Part1: part1,
	}
}

func day6Part2() Solution {
	races := day6Setup(true)
	winningStrats := make([]int, len(races))
	for i, r := range races {
		for t := 1; t < r.Time; t++ {
			time := t * (r.Time - t)
			if time > r.Distance {
				winningStrats[i]++
			}
		}
	}
	errorMargin := 1
	for _, s := range winningStrats {
		errorMargin *= s
	}
	return Solution{
		Part2: errorMargin,
	}
}
