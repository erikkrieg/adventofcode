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
	}
	return password
}

func (d *Day1) Part2() int {
	password := 0
	pos := 50
	for _, d := range d.data {
		dir := d[:1]
		dist, err := strconv.Atoi(d[1:])
		if err != nil {
			panic(err)
		}
		q, r := divmod(dist, 100)
		dist = r
		password += q
		if dir == "L" {
			dist *= -1
		}
		pos += dist
		if pos == 0 {
			password += 1
		} else if pos > 99 {
			pos %= 100
			password += 1
		} else if pos < 0 {
			pos += 100
			if pos-dist != 100 {
				password += 1
			}
		}
		fmt.Printf("%s %d %d %d\n", dir, dist, pos, password)
	}
	return password
}

func divmod(n, d int) (int, int) {
	q := n / d
	r := n % d
	return q, r
}

func init() {
	puzzleSolutions[0] = (&Day1{}).Solve
}
