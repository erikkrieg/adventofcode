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
	stacks, moves := parse_inputs()
	for _, m := range moves {
		doMove(stacks, &m)
	}
	fmt.Println(stacks, moves)
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

func parse_inputs() ([]Stack, []Move) {
	contents, err := ioutil.ReadFile("input/test-05.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputs := strings.Split(string(contents), "\n\n")
	stacks := parse_stacks(inputs[0])
	moves := parse_moves(inputs[1])
	return stacks, moves
}

func parse_stacks(input string) []Stack {
	// Mocking parsing of stacks to build movement logic first.
	return []Stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
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
