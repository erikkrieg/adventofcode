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
	day5BothParts().Print()
}

func day5Setup() ([]int, [][]string) {
	data := input.Lines("day-5")
	if useTestInput {
		data = input.Lines("test-5")
	}
	seeds := getNumbers(strings.Split(data[0], ": ")[1])
	categories := [][]string{}
	category := []string{}
	for i := 3; i < len(data); i++ {
		if data[i] == "" {
			categories = append(categories, category)
			category = []string{}
			i++
			continue
		}
		category = append(category, data[i])
	}
	categories = append(categories, category)
	return seeds, categories
}

func day5BothParts() Solution {
	seeds, categories := day5Setup()
	fmt.Println(seeds)

	closestLocation := math.MaxInt
	for _, seed := range seeds {
		location := getLocation(seed, categories)
		closestLocation = lib.Min(closestLocation, location)
	}

	// This is dumb brute force that will probs not work on real input.
	seeds = []int{
		2686125367, 200,
	}
	closestLocationPart2 := math.MaxInt
	nearestSeed := ""
	//limit := 4218385602
	limit := 2686125867
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println(i)
		if seeds[i] > limit {
			break
		}
		for r := 0; r < seeds[i+1]; r += 1 {
			if seeds[i]+r > limit {
				break
			}
			location := getLocation(seeds[i]+r, categories)
			if location < closestLocationPart2 {
				nearestSeed = fmt.Sprintf("Seed %d - i=%d, r=%d\n", seeds[i]+r, i, r)
			}
			closestLocationPart2 = lib.Min(closestLocationPart2, location)
		}
	}

	fmt.Println(nearestSeed)
	return Solution{
		Part1: closestLocation,
		Part2: closestLocationPart2,
	}
}

func getLocation(seed int, categories [][]string) int {
	number := seed
	for _, category := range categories {
		nextNumber := number
		for _, c := range category {
			numbers := getNumbers(c)
			destination := numbers[0]
			source := numbers[1]
			distance := numbers[2]
			sourceLimit := source + (distance - 1)
			if nextNumber <= sourceLimit && nextNumber >= source {
				nextNumber = destination + (number - source)
				break
			}
		}
		number = nextNumber
	}
	return number
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
