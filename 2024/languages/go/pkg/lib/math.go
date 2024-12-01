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

// Greatest Common Denominator
func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// Lowest Common Multiple
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func IntsLCM(ints ...int) int {
	return LCM(ints[0], ints[1], ints[2:]...)
}
