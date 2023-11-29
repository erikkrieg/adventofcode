package solutions

import "fmt"

func init() {
	puzzleSolutions[{{sub (atoi .date) 1}}] = day{{.date}}Solution
}

func day{{.date}}Solution() {
	fmt.Println("Day {{.date}}")
}
