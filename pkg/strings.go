package pkg

// ConcatStrings takes in a variadic number of strs and combines them into a one string
func ConcatStrings(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}

	return result
}
