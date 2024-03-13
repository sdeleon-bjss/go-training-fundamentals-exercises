package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
	"sort"
)

// Exercise 8:
// - Create a program that: [Write File][Read File][I/O Package][I/O]
// - Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// - Reads a list of cities from the newly created file.
// - Displays the list of cities in alphabetical order.
func main() {
	citiesInput := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

	err := pkg.WriteStringsToFile(citiesInput, "cities.txt")
	if err != nil {
		return
	}

	citiesOutput, err := pkg.ReadStringsFromFile("cities.txt")
	if err != nil {
		return
	}

	sort.Strings(citiesOutput)

	fmt.Println("unsorted:", citiesInput)
	fmt.Println("sorted:", citiesOutput)
}
