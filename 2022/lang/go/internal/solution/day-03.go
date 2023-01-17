package solution

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day03() {
	fmt.Println("- Day 03")
	file, err := os.Open("input/day-03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOneAnswer := 0
	partTwoAnswer := 0
	group := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		// Part one
		mid := len(line) / 2
		char, err := firstCommonChar([]string{line[0:mid], line[mid:]})
		if err != nil {
			fmt.Println("Part one error:", err)
			break
		}
		partOneAnswer += priority(char)

		// Part two
		group = append(group, line)
		if len(group) == 3 {
			char, err := firstCommonChar(group)
			if err != nil {
				fmt.Println("Part two error:", err)
				break
			}
			partTwoAnswer += priority(char)
			group = nil
		}
	}

	fmt.Println("  - Part 1: ", partOneAnswer)
	fmt.Println("  - Part 2: ", partTwoAnswer)
}

func priority(r rune) int {
	prio := int(r) - 38
	if prio > 52 {
		prio -= 58
	}
	return prio
}

func firstCommonChar(arr []string) (rune, error) {
chars:
	for _, r := range arr[0] {
		for _, s := range arr[1:] {
			if !strings.ContainsRune(s, r) {
				continue chars
			}
		}
		return r, nil
	}
	return ' ', errors.New("no common character shared among strings")
}
