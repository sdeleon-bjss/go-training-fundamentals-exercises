package pkg

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	FirstRoll  int
	SecondRoll int
	Combined   int
}

func (d *Dice) Roll() int {
	return rand.Intn(6) + 1
}

func (d *Dice) SetCombined() int {
	return d.FirstRoll + d.SecondRoll
}

func (d *Dice) ProcessRoll() string {
	switch d.Combined {
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

func (d *Dice) PrintRoll(attempt int) {
	fmt.Printf(
		"[Roll attempt #%d] First roll: %d Second roll: %d, Combined for: (%d)\nResult = %s\n-------------------------------------------------------------------\n",
		attempt, d.FirstRoll, d.SecondRoll, d.Combined, d.ProcessRoll())
}
