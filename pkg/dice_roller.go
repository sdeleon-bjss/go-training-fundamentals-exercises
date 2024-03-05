package pkg

import "math/rand"

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
