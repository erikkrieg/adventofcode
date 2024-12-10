package solutions

import (
	"fmt"
	"strconv"
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
		if d.evaluate(testValue, numbers[0], numbers[1:], false) {
			validTestValueSum += testValue
		}
	}
	return validTestValueSum
}

func (d *Day7) Part2() int {
	validTestValueSum := 0
	for _, line := range d.data {
		splitLine := strings.Split(line, ": ")
		testValue := lib.Atoi(splitLine[0])
		numbers := lib.Ints(strings.Fields(splitLine[1]))
		if d.evaluate(testValue, numbers[0], numbers[1:], true) {
			validTestValueSum += testValue
		}
	}
	return validTestValueSum
}

func (d *Day7) evaluate(testValue, currentValue int, remainingNumbers []int, includeConcat bool) bool {
	if len(remainingNumbers) == 0 {
		return testValue == currentValue
	}
	if testValue < currentValue {
		// Based on the sets of numbers and available operations the currentValue
		// can only increase.
		return false
	}
	next := remainingNumbers[0]
	rest := remainingNumbers[1:]
	evalAdd := d.evaluate(testValue, currentValue+next, rest, includeConcat)
	evalMul := d.evaluate(testValue, currentValue*next, rest, includeConcat)
	if !includeConcat {
		return evalAdd || evalMul
	}
	nextConcatValue := lib.Atoi(strconv.Itoa(currentValue) + strconv.Itoa(next))
	evalConcat := d.evaluate(testValue, nextConcatValue, rest, includeConcat)
	return evalAdd || evalMul || evalConcat
}

func init() {
	puzzleSolutions[6] = (&Day7{}).Solve
}
