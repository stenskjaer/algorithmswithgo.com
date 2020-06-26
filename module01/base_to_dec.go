package module01

import (
	"errors"
	"fmt"
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
func BaseToDec(value string, base int) int {
	// Zero-indexed length
	length := len(value) - 1

	// Get the integer value of the first character
	integerValue, err := valueInCharset(string(value[0]))
	if err != nil {
		fmt.Println("Failure: ", err)
	}

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

// valueInCharset converts a string value to an integer
// value within the used charset. This is done by iterating
// the charset and returning the index of the value, which
// reflects its numerical value.
func valueInCharset(character string) (int, error) {
	var res int
	var err error
	for i, val := range charset {
		if fmt.Sprintf("%c", val) == character {
			return i, err
		}
	}
	return res, errors.New("the character was not found in the charset")
}
