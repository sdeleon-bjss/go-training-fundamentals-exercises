package main

import "github.com/sdeleon-bjss/pkg"

// Exercise 07:
// Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.
// Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
// - 7 and 11 are to be called NATURAL
// - 2 is called SNAKE-EYES-CRAPS
// - 3 and 12 is called LOSS-CRAPS
// - Any other combination is called NEUTRAL.
func main() {
	dice := pkg.Dice{}

	for i := 1; i <= 50; i++ {
		dice.FirstRoll = dice.Roll()
		dice.SecondRoll = dice.Roll()
		dice.Combined = dice.SetCombined()
		dice.PrintRoll(i)
	}
}
