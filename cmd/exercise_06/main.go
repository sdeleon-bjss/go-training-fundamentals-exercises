package main

import (
	"bjss-go-training/pkg"
	"time"
)

// Exercise 6:
// Create a program that calculates the age of a person given their date of birth. [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
// (Use the github.com/bearbin/go-age to aid in the creation of this app. Also review unit testing applied to the application age.go within the imported package.)
func main() {
	people := []pkg.Person{
		{DOB: time.Date(1992, time.January, 28, 0, 0, 0, 0, time.UTC)},
		{DOB: time.Date(1992, time.August, 31, 0, 0, 0, 0, time.UTC)},
		{DOB: time.Date(1975, time.May, 21, 0, 0, 0, 0, time.UTC)},
		{DOB: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC)},
		{DOB: time.Date(2023, time.February, 10, 0, 0, 0, 0, time.UTC)},
	}

	for _, person := range people {
		println(person.Age())
	}
}
