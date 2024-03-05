package main

import (
	"fmt"
	"github.com/bearbin/go-age"
	"time"
)

// Create a program that calculates the age of a person given their date of birth. [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
// (Use the github.com/bearbin/go-age to aid in the creation of this app. Also review unit testing applied to the application age.go within the imported package.)
func main() {
	dob := time.Date(1992, time.January, 28, 9, 15, 0, 0, time.UTC)
	ageResult := CalculateAge(dob)

	fmt.Println("You are", ageResult, "years old")
}

func CalculateAge(dob time.Time) int {
	ageInYears := age.Age(dob)
	return ageInYears
}
