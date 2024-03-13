package main

import "github.com/sdeleon-bjss/pkg"

// Exercise 1:
// Create a program that has multiple string variable and displays the string on one line. [Strings]
func main() {
	messageOne := "Hello, my name is"
	messageTwo := "Steven and I am finally using Go!"

	output := pkg.ConcatStrings(messageOne, " ", messageTwo)

	println(output)
}
