package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

var day2Test = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func init() {
	puzzleSolutions[1] = day2Solution
}

func day2Solution() {
	fmt.Println("Day 2")
	fmt.Printf(" Part 1: %d\n", day2Part1(&day2Test))
	fmt.Printf(" Part 2: %d\n", day2Part2(&day2Test))
}

func day2Part1(test *[]string) int {
	games := input.Lines("day-2")
	if useTestInput {
		games = *test
	}
	sum := 0
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for i, game := range games {
		id := i + 1
		trimmedGame := game[strings.IndexRune(game, ':')+1:]
		sets := strings.Split(trimmedGame, ";")
	rangeSets:
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				c := strings.Split(strings.TrimSpace(cube), " ")
				color := c[1]
				count, err := strconv.Atoi(c[0])
				if err != nil {
					panic(err)
				}
				if limits[color] < count {
					id = 0
					break rangeSets
				}
			}
		}
		sum += id
	}
	return sum
}

func day2Part2(test *[]string) int {
	games := input.Lines("day-2")
	if useTestInput {
		games = *test
	}
	minimumCubeSetPowerSum := 0
	for _, game := range games {
		minCubeSet := make(map[string]int)
		trimmedGame := game[strings.IndexRune(game, ':')+1:]
		sets := strings.Split(trimmedGame, ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				c := strings.Split(strings.TrimSpace(cube), " ")
				color := c[1]
				count, err := strconv.Atoi(c[0])
				if err != nil {
					panic(err)
				}
				if count > minCubeSet[color] {
					minCubeSet[color] = count
				}
			}
		}
		power := 1
		for _, count := range minCubeSet {
			power *= count
		}
		minimumCubeSetPowerSum += power
	}
	return minimumCubeSetPowerSum
}
