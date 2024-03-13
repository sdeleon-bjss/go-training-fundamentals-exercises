package main

import "github.com/sdeleon-bjss/pkg"

// Exercise 10:
// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]
func main() {
	var pupils []pkg.Student

	// grabbed names from
	// - https://www.scarymommy.com/pregnancy/first-middle-names-boys-girls
	// - https://www.thoughtco.com/most-common-us-surnames-1422656
	student1, _ := pkg.NewPupil("Albert", "Blake", "Smith", "03/14/2012")
	student2, _ := pkg.NewPupil("Archie", "James", "Johnson", "04/15/2013")
	student3, _ := pkg.NewPupil("Benjamin ", "Tate", "Williams", "05/16/2014")
	student4, _ := pkg.NewPupil("Benson", "Scott", "Brown", "06/17/2015")
	student5, _ := pkg.NewPupil("Blaine", "Andrew", "Jones", "07/18/2016")
	student6, _ := pkg.NewPupil("Caleb", "Cole", "Garcia", "08/19/2017")
	student7, _ := pkg.NewPupil("Carter", "Ezra", "Miller", "09/20/2018")
	student8, _ := pkg.NewPupil("Charles", "Rupert", "Davis", "10/21/2019")
	student9, _ := pkg.NewPupil("Cooper", "Nolan", "Rodriguez", "11/22/2020")
	student10, _ := pkg.NewPupil("David", "Michael", "Martinez", "12/23/2021")

	pupils = append(pupils, student1, student2, student3, student4, student5, student6, student7, student8, student9, student10)

	for _, pupil := range pupils {
		println(pupil.GetFullName(), pupil.GetDateOfBirth(), pupil.GetAge())
	}
}
