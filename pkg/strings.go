package pkg

func ConcatStrings(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}

	return result
}
