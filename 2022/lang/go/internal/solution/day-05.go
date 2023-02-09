package solution

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day05() {
	fmt.Println("- Day 05")
	stacks, moves := parse_inputs("input/day-05.txt")
	for _, m := range moves {
		doMove(stacks, &m)
	}
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
	// Empty line for cleaner stdout
	fmt.Println()
}

type Move struct {
	quantity int
	from     int
	to       int
}
type Crate = string
type Stack = []Crate

func doMove(s []Stack, m *Move) {
	for i := 0; i < m.quantity; i++ {
		end := len(s[m.from]) - 1
		s[m.to] = append(s[m.to], s[m.from][end])
		s[m.from] = s[m.from][:end]
	}
}

func parse_inputs(path string) ([]Stack, []Move) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	inputs := strings.Split(string(contents), "\n\n")
	stacks := parse_stacks(inputs[0])
	moves := parse_moves(inputs[1])
	return stacks, moves
}

func parse_stacks(input string) []Stack {
	re := regexp.MustCompile("[A-Z]{1}")
	lines := strings.Split(input, "\n")
	stacks := make([]Stack, (len(lines[0])+1)/4)
	for _, line := range lines {
		for i := 1; i < len(line)-1; i += 4 {
			crate := line[i : i+1]
			stackIndex := i / 4
			if re.MatchString(crate) {
				stacks[stackIndex] = append(Stack{crate}, stacks[stackIndex]...)
			}
		}
	}
	return stacks
}

func parse_moves(input string) []Move {
	re := regexp.MustCompile("[0-9]+")
	moves := []Move{}
	for _, m := range strings.Split(input, "\n") {
		fields := [3]int{}
		for i, v := range re.FindAllString(m, -1) {
			v, _ := strconv.Atoi(v)
			fields[i] = v
		}
		moves = append(moves, Move{quantity: fields[0], from: fields[1] - 1, to: fields[2] - 1})
	}
	return moves
}
