package main

import "bjss-go-training/pkg"

// Exercise 10:
// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]
func main() {
	var pupils []pkg.Person

	pupils = append(pupils, pkg.NewPupil("Steven", "William", "DeLeon", "01/28/1992", 32))
	pupils = append(pupils, pkg.NewPupil("Gracie", "my", "cat", "06/22/2018", 6))
	pupils = append(pupils, pkg.NewPupil("Penelope", "my", "cat", "05/12/2017", 7))
	pupils = append(pupils, pkg.NewPupil("Mr. Boot", "my", "cat", "10/18/2019", 5))

	for _, pupil := range pupils {
		println(pupil.GetFullName())
		println(pupil.GetDateOfBirth())
		println(pupil.GetAge())
	}
}
