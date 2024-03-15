package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 2:
// Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line. [Keyboard input]
func main() {
	firstName := getInput("Enter your first name: ")
	middleName := getInput("Enter your middle name: ")
	lastName := getInput("Enter your last name: ")

	name := pkg.Name{}
	err := name.SetFullName(pkg.GetFullName(firstName, middleName, lastName))
	if err != nil {
		println(err)
		return
	}

	fullName := name.GetFullName()
	println("Your full name is: ", fullName)
}

func getInput(prompt string) string {
	var input string
	fmt.Println(prompt)

	_, err := fmt.Scan(&input)
	if err != nil {
		return ""
	}

	return input
}
