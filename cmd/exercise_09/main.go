package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 9:
// Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]
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

	nameSliced := name.SliceName()
	name.Print(nameSliced)
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
