package lib

import "strconv"

func Map[In any, Out any](input []In, mapper func(In) Out) []Out {
	out := make([]Out, len(input))
	for i, v := range input {
		out[i] = mapper(v)
	}
	return out
}

// Ints will attempt to convert a slice of strings to a slice of ints.
// If an error is encountered it will panic.
func Ints(in []string) []int {
	return Map(in, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	})
}
