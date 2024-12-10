package solutions

import (
	"fmt"
	"strings"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day7 struct {
	data []string
}

func (d *Day7) Setup() {
	data := input.Lines("day-7")
	if useTestInput {
		data = input.Lines("test-7")
	}
	d.data = data
}

func (d *Day7) Solve() {
	fmt.Println("Day 7")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day7) Part1() int {
	validTestValueSum := 0
	for _, line := range d.data {
		splitLine := strings.Split(line, ": ")
		testValue := lib.Atoi(splitLine[0])
		numbers := lib.Ints(strings.Fields(splitLine[1]))
		if d.evaluate(testValue, numbers[0], numbers[1:]) {
			validTestValueSum += testValue
		}
	}
	return validTestValueSum
}

func (d *Day7) evaluate(testValue, currentValue int, remainingNumbers []int) bool {
	if len(remainingNumbers) == 0 {
		return testValue == currentValue
	}
	if testValue < currentValue {
		// Based on the sets of numbers and available operations the currentValue
		// can only increase.
		return false
	}
	next := remainingNumbers[0]
	nextAdd := currentValue + next
	nextMul := currentValue * next
	rest := remainingNumbers[1:]
	return d.evaluate(testValue, nextAdd, rest) || d.evaluate(testValue, nextMul, rest)

}

func (d *Day7) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[6] = (&Day7{}).Solve
}
