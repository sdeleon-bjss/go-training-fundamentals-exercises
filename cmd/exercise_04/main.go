package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 4:
// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// - Display the array content in ascending sequential order 1 to 10.
// - Display the array content in descending sequential order 10 to 1.
// - Count even numbers and odd numbers in increasing and decreasing sequential order.
// - Display the even and odd count sequences to screen.
func main() {
	var numbers []int
	for i := range 10 {
		numbers = append(numbers, i+1)
	}

	ascending := pkg.Ascending(numbers)
	descending := pkg.Descending(numbers)

	fmt.Println("Ascending:", ascending)
	fmt.Println("Descending:", descending)

	oddAscending, evenAscending := pkg.OddEvenCount(ascending, "ascending")
	fmt.Println("Odd numbers in ascending order:", oddAscending)
	fmt.Println("Even numbers in ascending order:", evenAscending)

	oddDescending, evenDescending := pkg.OddEvenCount(descending, "descending")
	fmt.Println("Odd numbers in descending order:", oddDescending)
	fmt.Println("Even numbers in descending order:", evenDescending)
}
