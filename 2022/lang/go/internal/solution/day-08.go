package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Void struct{}

var void Void

func Day08() {
	fmt.Println("- Day 08")
	file, err := os.Open("input/day-08.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	trees := parseTrees(scanner)
	partOne(trees)
	partTwo(trees)
}

func partOne(trees [][]int) {
	visibleTrees := make(map[string]Void)

	for y := range trees {
		viewTreesHorizontally(y, trees, visibleTrees, false)
		viewTreesHorizontally(y, trees, visibleTrees, true)
	}

	for x := range trees[0] {
		viewTreesVertically(x, trees, visibleTrees, false)
		viewTreesVertically(x, trees, visibleTrees, true)
	}

	fmt.Println("  - Part 1: ", len(visibleTrees))
}

type Point struct {
	x int
	y int
}

func (p *Point) shift(target *Point) Point {
	x := p.x
	if x > target.x {
		x -= 1
	} else if x < target.x {
		x += 1
	}
	y := p.y
	if y > target.y {
		y -= 1
	} else if y < target.y {
		y += 1
	}

	return Point{x, y}
}

func sightLine(a Point, b Point, trees [][]int) int {
	cursor := a
	cmp := trees[a.y][a.x]
	line := []int{}
	for cursor != b {
		cursor = cursor.shift(&b)
		tree := trees[cursor.y][cursor.x]
		line = append(line, tree)
		if tree >= cmp {
			break
		}
	}
	return len(line)
}

func partTwo(trees [][]int) {
	maxScenicScore := 0
	for y, row := range trees {
		for x := range row {
			lineUp := sightLine(Point{x, y}, Point{x, 0}, trees)
			lineRight := sightLine(Point{x, y}, Point{len(row) - 1, y}, trees)
			lineDown := sightLine(Point{x, y}, Point{x, len(trees) - 1}, trees)
			lineLeft := sightLine(Point{x, y}, Point{0, y}, trees)
			scenicScore := lineUp * lineDown * lineRight * lineLeft
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	fmt.Println("  - Part 2: ", maxScenicScore)
}

func viewTreesHorizontally(rowIndex int, trees [][]int, visibleTrees map[string]Void, reverse bool) {
	viewed := []int{}
	row := trees[rowIndex]
	for x := range row {
		if reverse {
			x = len(row) - (x + 1)
		}
		tree := row[x]
		if len(viewed) == 0 || tree > viewed[len(viewed)-1] {
			viewed = append(viewed, tree)
			k := fmt.Sprintf("%d-%d", x, rowIndex)
			if _, found := visibleTrees[k]; !found {
				visibleTrees[k] = void
			}
		}
	}
}

func viewTreesVertically(columnIndex int, trees [][]int, visibleTrees map[string]Void, reverse bool) {
	viewed := []int{}
	for y := range trees {
		if reverse {
			y = len(trees) - (y + 1)
		}
		tree := trees[y][columnIndex]
		if len(viewed) == 0 || tree > viewed[len(viewed)-1] {
			viewed = append(viewed, tree)
			k := fmt.Sprintf("%d-%d", columnIndex, y)
			if _, found := visibleTrees[k]; !found {
				visibleTrees[k] = void
			}
		}
	}
}

func parseTrees(input *bufio.Scanner) [][]int {
	trees := [][]int{}
	for input.Scan() {
		line := input.Text()
		treeRow := []int{}
		for _, char := range line {
			treeRow = append(treeRow, int(char-'0'))
		}
		trees = append(trees, treeRow)
	}
	return trees
}
