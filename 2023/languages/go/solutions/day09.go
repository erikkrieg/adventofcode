package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[8] = day9Solution
}

func setupDay9() []string {
	data := input.Lines("day-9")
	if useTestInput {
		data = input.Lines("test-9.1")
	}
	return data
}

func day9Solution() {
	fmt.Println("Day 9")
	data := setupDay9()
	Solution{
		Part1: day9part1(data),
		Part2: day9part2(data),
	}.Print()
}

func parseValues(data string) []int {
	values := []int{}
	for _, d := range strings.Split(data, " ") {
		v, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		values = append(values, v)
	}
	return values
}

func day9part1(data []string) int {
	nextValues := make([]int, len(data))
	for i, d := range data {
		values := parseValues(d)
		nextValues[i] = calculateNextValue(values)
	}
	sum := 0
	for _, v := range nextValues {
		sum += v
	}
	return sum
}

func day9part2(data []string) int {
	prevValues := make([]int, len(data))
	for i, d := range data {
		values := parseValues(d)
		prevValues[i] = calculatePrevValue(values)
	}
	sum := 0
	for _, v := range prevValues {
		sum += v
	}
	return sum
}

func calculateNextValue(values []int) int {
	differences := calcDiffs(values)
	for i := len(differences) - 1; i > 0; i-- {
		last := len(differences[i]) - 1
		v := differences[i][last]
		nlast := len(differences[i-1]) - 1
		nv := differences[i-1][nlast]
		differences[i-1] = append(differences[i-1], v+nv)
	}
	return differences[0][len(differences[0])-1]
}

func calculatePrevValue(values []int) int {
	differences := calcDiffs(values)
	for i := len(differences) - 1; i > 0; i-- {
		v := differences[i][0]
		pv := differences[i-1][0]
		differences[i-1] = append([]int{pv - v}, differences[i-1]...)
	}
	return differences[0][0]
}

func calcDiffs(values []int) [][]int {
	differences := make([][]int, 0)
	differences = append(differences, values)
	for {
		d, done := calcDiff(differences[len(differences)-1])
		differences = append(differences, d)
		if done {
			break
		}
	}
	return differences
}

func calcDiff(values []int) ([]int, bool) {
	zeroCount := 0
	d := make([]int, len(values)-1)
	for i, v := range values {
		if i == len(values)-1 {
			break
		}
		d[i] = values[i+1] - v
		if d[i] == 0 {
			zeroCount++
		}
	}
	return d, zeroCount == len(d)
}
