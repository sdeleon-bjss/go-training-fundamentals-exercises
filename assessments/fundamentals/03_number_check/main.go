package main

import "fmt"

// Create a program that allows the user to input a number. Check whether the number lies between 1 and 10. [Variables]
func main() {
	var inputNumber int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&inputNumber)
	
	number := IsNumberBetweenRange(inputNumber, 1, 10)

	println("Is the number between 1 and 10? ", number)
}

func IsNumberBetweenRange(num int, start int, end int) bool {
	if num >= start && num <= end {
		return true
	}

	return false
}
