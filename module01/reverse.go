package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var result string

	// Go through the string from the beginning
	for _, char := range word {
		// Build a new string, adding the current char _before_ the result so far.
		result = string(char) + result
	}

	return result
}
