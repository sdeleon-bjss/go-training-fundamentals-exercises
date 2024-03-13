package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 9:
// Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]
func main() {
	fullName := "Steven William DeLeon"

	name := pkg.Name{}
	err := name.SetFullName(fullName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("full-name : %s, middle-name : %s, and surname : %s ", name.GetFullName(), name.Middle, name.Last)
}
