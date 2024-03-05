package main

import "bjss-go-training/pkg"

// Exercise 10:
// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]
func main() {
	pupils := []pkg.Pupil{
		pkg.NewPupil("Steven", "William", "DeLeon", "01/28/1992", 32),
		pkg.NewPupil("Gracie", "my", "cat", "06/22/2018", 6),
		pkg.NewPupil("Penelope", "my", "cat", "05/12/2017", 7),
		pkg.NewPupil("Mr. Boot", "my", "cat", "10/18/2019", 5),
		// ... this feels like a good place to stop :)
	}

	for _, pupil := range pupils {
		println(pupil.GetFullName())
		println(pupil.DateOfBirth)
		println(pupil.Age)
		println()
	}
}
