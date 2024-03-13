package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 2:
// Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line. [Keyboard input]
func main() {
	var firstName, middleName, lastName string
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your middle name: ")
	fmt.Scan(&middleName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fullName := pkg.GetFullName(firstName, middleName, lastName)

	println("Your full name is: ", fullName)
}
