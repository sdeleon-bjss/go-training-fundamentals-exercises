package exercise_01_strings

func ConcatStrings(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}

	return result
}
