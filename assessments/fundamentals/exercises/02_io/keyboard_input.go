package main

import "fmt"

func main() {
	var firstName, middleName, lastName string
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your middle name: ")
	fmt.Scan(&middleName)
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Your full name is: %s %s %s\n", firstName, middleName, lastName)
}
