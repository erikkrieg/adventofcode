package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

var day3Test = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func init() {
	puzzleSolutions[2] = day3Solution
}

func day3Solution() {
	fmt.Println("Day 3")
	fmt.Printf(" - Part 1: %d\n", day3Part1(day3Test))
}

type Edge struct {
	y     int
	x     int
	chars string
}

func day3Part1(schematic []string) int {
	digits := "0123456789"
	if !useTestInput {
		schematic = input.Lines("day-3")
	}
	partNumSum := 0
	gears := make(map[int]map[int][]int)
	for y, s := range schematic {
		num := ""
		for x, char := range s {
			isDigit := strings.ContainsRune(digits, char)
			if isDigit {
				num += string(char)
			}
			if (!isDigit || x == len(s)-1) && num != "" {
				isPart := false
				xmin := lib.Max(x-len(num)-1, 0)
				xmax := lib.Min(x+1, len(s))
				edges := []Edge{}
				if y > 0 {
					edge := Edge{y: y - 1, x: xmin, chars: schematic[y-1][xmin:xmax]}
					edges = append(edges, edge)
				}
				if y < len(schematic)-1 {
					edge := Edge{y: y + 1, x: xmin, chars: schematic[y+1][xmin:xmax]}
					edges = append(edges, edge)
				}
				if xmin-1 > 0 {
					edge := Edge{y: y, x: xmin, chars: s[xmin : xmin+1]}
					edges = append(edges, edge)
				}
				if xmax <= len(s) {
					edge := Edge{y: y, x: x, chars: s[x:xmax]}
					edges = append(edges, edge)
				}
				n, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				for _, e := range edges {
					side := e.chars
					if strings.ContainsAny(side, "!@#$%^&*-=+/:") {
						gearX := strings.IndexRune(side, '*')
						if gearX > -1 {
							gearX += e.x
							if _, ok := gears[e.y]; !ok {
								gears[e.y] = make(map[int][]int)
							}
							gears[e.y][gearX] = append(gears[e.y][gearX], n)
						}
						isPart = true
						break
					}
				}
				if isPart {
					partNumSum += n
				}
				num = ""
			}

		}
	}
	gearRatioSum := 0
	for _, row := range gears {
		for _, gear := range row {
			if len(gear) == 2 {
				gearRatioSum += gear[0] * gear[1]
			}
		}
	}
	// TODO: refactor so that part 1 and 2 can both be executed
	// return partNumSum
	return gearRatioSum
}
