package solutions

import "fmt"

func init() {
	puzzleSolutions[{{.date}}] = day{{.date}}Solution
}

func day{{.date}}Solution() {
	fmt.Println("Day {{.date}}")
}
