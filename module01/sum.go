package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	result := 0
	for _, value := range numbers {
		result += value
	}
	return result
}
