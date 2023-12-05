package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[4] = day5Solution
}

func day5Solution() {
	fmt.Println("Day 5")
	day5Part1().Print()
}

func day5Part1() Solution {
	data := input.Lines("day-5")
	if useTestInput {
		data = input.Lines("test-5")
	}
	closestLocation := math.MaxInt
	seeds := getNumbers(strings.Split(data[0], ": ")[1])
	for _, seed := range seeds {
		location := getLocation(seed, data[1:])
		closestLocation = lib.Min(closestLocation, location)
	}

	return Solution{
		Part1: closestLocation,
	}
}

func getLocation(seed int, data []string) int {
	fmt.Printf("Seed: %d\n", seed)
	number := seed
	nextNumber := number
	for i := 0; i < len(data); i++ {
		row := data[i]
		if row == "" {
			fmt.Printf("  - Next row: %s\n", data[i+1])
			i++
			number = nextNumber
			continue
		}

		numbers := getNumbers(row)
		destination := numbers[0]
		source := numbers[1]
		distance := numbers[2]

		sourceLimit := source + (distance - 1)
		if number <= sourceLimit && number >= source {
			nextNumber = destination + (number - source)
			fmt.Printf("    - Number changed from %d to %d\n", number, nextNumber)
		}
	}
	return nextNumber
}

func getNumbers(data string) []int {
	nums := []int{}
	for _, s := range strings.Split(data, " ") {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
