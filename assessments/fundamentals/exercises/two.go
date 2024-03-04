package exercises

import (
	"fmt"
)

func GetUsersFullName() string {
	var firstName, middleName, lastName string
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your middle name")
	fmt.Scan(&middleName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fullName := fmt.Sprintf("%s %s %s", firstName, middleName, lastName)

	return fullName
}
