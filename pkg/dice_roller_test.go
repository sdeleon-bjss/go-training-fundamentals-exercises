package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
)

func TestProcessRoll(t *testing.T) {
	testCases := []struct {
		roll int
		want string
	}{
		{roll: 1, want: "NEUTRAL"},
		{roll: 2, want: "SNAKE-EYES-CRAPS"},
		{roll: 3, want: "LOSS-CRAPS"},
		{roll: 4, want: "NEUTRAL"},
		{roll: 5, want: "NEUTRAL"},
		{roll: 6, want: "NEUTRAL"},
		{roll: 7, want: "NATURAL"},
		{roll: 8, want: "NEUTRAL"},
		{roll: 9, want: "NEUTRAL"},
		{roll: 10, want: "NEUTRAL"},
		{roll: 11, want: "NATURAL"},
		{roll: 12, want: "LOSS-CRAPS"},
	}

	for _, test := range testCases {
		result := pkg.ProcessRoll(test.roll)
		if result != test.want {
			t.Errorf("want %s but got %s", test.want, result)
		}
	}
}

func TestRollDice(t *testing.T) {
	diceRoll := pkg.RollDice()

	if diceRoll < 1 || diceRoll > 6 {
		t.Errorf("want a number between 1 and 6 but got %d", diceRoll)
	}
}

func TestComboRolls(t *testing.T) {
	result := pkg.ComboRolls()

	if result[0] < 1 || result[0] > 6 {
		t.Errorf("want a number between 1 and 6 but got %d", result[0])
	}

	if result[1] < 1 || result[1] > 6 {
		t.Errorf("want a number between 1 and 6 but got %d", result[1])
	}
}
