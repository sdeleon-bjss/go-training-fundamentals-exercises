package main

import (
	"fmt"
	"math/rand"
)

// Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.
// Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
// - 7 and 11 are to be called NATURAL
// - 2 is called SNAKE-EYES-CRAPS
// - 3 and 12 is called LOSS-CRAPS
// - Any other combination is called NEUTRAL.
func main() {
	for i := 0; i < 50; i++ {
		rolls := ComboRolls()
		combinedRolls := rolls[0] + rolls[1]
		rollResult := ProcessRoll(combinedRolls)
		output := fmt.Sprintf("Roll attempt #%d: You rolled a (%d) = %s", i+1, combinedRolls, rollResult)
		fmt.Println(output)
	}
}

func ComboRolls() [2]int {
	return [2]int{RollDice(), RollDice()}
}

func RollDice() int {
	return rand.Intn(6) + 1
}

func ProcessRoll(roll int) string {
	switch roll {
	case 7, 11:
		return "NATURAL"
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 3, 12:
		return "LOSS-CRAPS"
	default:
		return "NEUTRAL"
	}
}
