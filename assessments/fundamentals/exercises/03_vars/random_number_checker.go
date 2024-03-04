package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberInput string
	fmt.Println("Please enter a random number between 1 and 10")
	fmt.Scan(&numberInput)

	parsedNumber, err := strconv.Atoi(numberInput)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	isBetween := isNumberBetweenRange(parsedNumber, 1, 10)

	fmt.Println("Is between 1 and 10: ", isBetween)
}

func isNumberBetweenRange(num int, start int, end int) bool {
	if num >= start && num <= end {
		return true
	}

	return false
}
