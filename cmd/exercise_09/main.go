package main

import (
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 9:
// Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]
func main() {
	fullName := "Steven William DeLeon"

	name := pkg.Name{}
	_ = name.SetFullName(fullName)

	name.Print()
}
