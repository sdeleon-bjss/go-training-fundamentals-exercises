package main

import "bjss-go-training/pkg"

// Exercise 9:
// Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]
func main() {
	firstName := "Steven"
	middleName := "William"
	lastName := "DeLeon"

	fullName := pkg.GetFullName(firstName, middleName, lastName)

	println(fullName)

	name := pkg.Name{
		First:  fullName[:len(firstName)],
		Middle: fullName[len(firstName)+1 : len(firstName)+1+len(middleName)],
		Last:   fullName[len(firstName)+1+len(middleName)+1:],
	}

	println(name.GetFullName())
}
