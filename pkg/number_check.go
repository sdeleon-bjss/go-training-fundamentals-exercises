package pkg

func IsNumberBetweenRange(num int, start int, end int) bool {
	if num >= start && num <= end {
		return true
	}

	return false
}
