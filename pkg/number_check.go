package pkg

// IsNumberBetweenRange checks if a number is between a given range
func IsNumberBetweenRange(num int, start int, end int) bool {
	if num >= start && num <= end {
		return true
	}

	return false
}
