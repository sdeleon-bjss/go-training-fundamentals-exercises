package pkg

import "math/rand"

// ComboRolls returns a pair of dice rolls
func ComboRolls() [2]int {
	return [2]int{RollDice(), RollDice()}
}

// RollDice returns a random number between 1 and 6
func RollDice() int {
	return rand.Intn(6) + 1
}

// ProcessRoll returns the outcome of the roll
//
// # Note - the parameters expect to receive the sum of the two dice rolls
//
// - 7 and 11 are to be called NATURAL
//
// - 2 is called SNAKE-EYES-CRAPS
//
// - 3 and 12 is called LOSS-CRAPS
//
// - Any other combination is called NEUTRAL.
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
