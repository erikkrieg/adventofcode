package lib

import (
	"strconv"
)

// Atoi quick and dirty. Handles errors by panicking.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
