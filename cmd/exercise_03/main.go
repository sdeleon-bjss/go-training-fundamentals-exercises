package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 3:
// Create a program that allows the user to input a number. Check whether the number lies between 1 and 10. [Variables]
func main() {
	var inputNumber int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&inputNumber)

	result := pkg.IsNumberBetweenRange(inputNumber, 1, 10)

	println("Is the number between 1 and 10? ", result)
}
