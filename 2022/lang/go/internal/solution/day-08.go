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

	fmt.Println(len(visibleTrees))
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
