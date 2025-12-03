package solutions

import (
	"fmt"
	"strconv"

	"github.com/erikkrieg/adventofcode/2025/pkg/input"
)

type Day1 struct {
	data []string
}

func (d *Day1) Setup() {
	data := input.Lines("day-1")
	if useTestInput {
		data = input.Lines("test-1")
	}
	d.data = data
}

func (d *Day1) Solve() {
	fmt.Println("Day 1")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day1) Part1() int {
	password := 0
	pos := 50
	for _, d := range d.data {
		dir := d[:1]
		dist, err := strconv.Atoi(d[1:])
		if err != nil {
			panic(err)
		}
		if dir == "L" {
			dist *= -1
		}
		pos += dist
		// Distance can be > 100, but mod on the pos handles case for both positive
		// and negative distances.
		pos %= 100
		if pos < 0 {
			pos += 100
		}
		if pos == 0 {
			password += 1
		}
		fmt.Printf("%s %d = %d\n", dir, dist, pos)
	}
	return password
}

func (d *Day1) Part2() int {
	return 0
}

func init() {
	puzzleSolutions[0] = (&Day1{}).Solve
}
