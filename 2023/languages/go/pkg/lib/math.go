package lib

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	return Max(n, -n)
}
