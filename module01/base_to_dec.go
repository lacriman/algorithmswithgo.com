package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	decNum := 0
	for i := 0; i < len(value); i++ {
		pos := len(value) - 1 - i
		digit := strings.IndexRune(charset, rune(value[i]))
		digit = digit * powInt(base, pos)
		decNum += int(digit)
	}
	return decNum
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
