package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
const charset = "0123456789ABCDEF"

func DecToBase(dec, base int) string {
	// Get remainder from division between decimal and base
	remainder := string(charset[dec%base])

	// Divide by base
	divided := dec / base

	// If division remainder is more than 0, repeat
	if divided > 0 {
		// Add recursive results to current remainder
		return DecToBase(divided, base) + remainder
	}

	return remainder
}

// DecToBaseAlternative is a more procedural alternative
// with a loop mutating the decimal value and adding to the
// result string for each iteration.
func DecToBaseAlternative(dec, base int) string {
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec /= base
	}
	return res
}
