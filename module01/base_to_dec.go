package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
// Method: The value is multiplied with the base raised to
// the power of the position in value string, which reflects
// its weight.
//
// This solution takes a non-recursive approach. I find this
// to be simpler to read, but grokking how the exponents are
// calculated takes a bit more concentration.
func BaseToDec(value string, base int) int {
	var result int
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		digit := charToDigit[string(value[i])]

		result += digit * multiplier
		multiplier = multiplier * base
	}
	return result
}

// BaseToDecAlternative presents an alternative
// recursion-based solution. The calculation of the
// exponents is more explicit here, and I think that is an
// advantage.
func BaseToDecAlternative(value string, base int) int {
	// Zero-indexed length
	length := len(value) - 1

	// Get the integer value of the first character
	integerValue := charToDigit[string(value[0])]

	// Calculate the conversion of the current character by
	// multiplying its integer value with the base value
	// raised to the power of the length of the value string
	// (indicating the weight of the current value).
	result := integerValue * int(math.Pow(float64(base), float64(length)))

	// Keep going until we are through the full value string
	if length > 0 {
		return BaseToDec(value[1:], base) + result
	}
	return result
}

var charToDigit = map[string]int {
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}
