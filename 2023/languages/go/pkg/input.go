package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Scan(puzzle string) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("input/%s.txt", puzzle))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return bufio.NewScanner(file)
}

func Chars(puzzle string) []rune {
	var chars []rune
	scanner := Scan(puzzle)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			chars = append(chars, char)
		}
	}
	return chars
}
