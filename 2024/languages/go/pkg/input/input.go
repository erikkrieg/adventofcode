package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Scan(puzzle string) (*bufio.Scanner, *os.File) {
	filename := fmt.Sprintf("inputs/%s.txt", puzzle)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file), file
}

func Chars(puzzle string) []rune {
	var chars []rune
	scanner, file := Scan(puzzle)
	defer file.Close()
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			chars = append(chars, char)
		}
	}
	return chars
}

func Lines(puzzle string) []string {
	var lines []string
	scanner, file := Scan(puzzle)
	defer file.Close()
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func LinesChars(puzzle string) [][]rune {
	var linesChars [][]rune
	scanner, file := Scan(puzzle)
	defer file.Close()
	for scanner.Scan() {
		var chars []rune
		for _, char := range scanner.Text() {
			chars = append(chars, char)
		}
		linesChars = append(linesChars, chars)
	}
	return linesChars
}
