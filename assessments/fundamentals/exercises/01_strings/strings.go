package main

import "fmt"

func main() {
	stringOne := "Hello there"
	stringTwo := "my name is"
	name := "Steven!"

	strings := []string{"I", "am", "finally", "using", "go"}

	combinedOutput := fmt.Sprintf("%s, %s %s", stringOne, stringTwo, name)

	var combinedOutput2 string

	for i, str := range strings {
		if i == 0 {
			combinedOutput2 += str
			continue
		}

		combinedOutput2 += " " + str
	}

	// outputs
	fmt.Println(combinedOutput)
	fmt.Println(combinedOutput2)
}
