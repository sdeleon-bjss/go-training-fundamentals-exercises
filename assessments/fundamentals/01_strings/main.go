package main

// Create a program that has multiple string variable and displays the string on one line. [Strings]
func main() {
	messageOne := "Hello, my name is"
	messageTwo := "Steven and I am finally using Go!"

	output := ConcatStrings(messageOne, " ", messageTwo)

	println(output)
}

func ConcatStrings(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}

	return result
}
