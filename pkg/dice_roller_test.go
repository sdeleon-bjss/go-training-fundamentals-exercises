package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testRoll struct {
	roll     int
	expected string
}

func TestProcessRoll(t *testing.T) {
	testCases := []testRoll{
		{roll: 1, expected: "NEUTRAL"},
		{roll: 2, expected: "SNAKE-EYES-CRAPS"},
		{roll: 3, expected: "LOSS-CRAPS"},
		{roll: 4, expected: "NEUTRAL"},
		{roll: 5, expected: "NEUTRAL"},
		{roll: 6, expected: "NEUTRAL"},
		{roll: 7, expected: "NATURAL"},
		{roll: 8, expected: "NEUTRAL"},
		{roll: 9, expected: "NEUTRAL"},
		{roll: 10, expected: "NEUTRAL"},
		{roll: 11, expected: "NATURAL"},
		{roll: 12, expected: "LOSS-CRAPS"},
	}

	for _, test := range testCases {
		result := pkg.ProcessRoll(test.roll)
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}

func TestRollDice(t *testing.T) {
	diceRoll := pkg.RollDice()

	if diceRoll < 1 || diceRoll > 6 {
		t.Errorf("Expected a number between 1 and 6 but got %d", diceRoll)
	}
}

func TestComboRolls(t *testing.T) {
	result := pkg.ComboRolls()

	if result[0] < 1 || result[0] > 6 {
		t.Errorf("Expected a number between 1 and 6 but got %d", result[0])
	}

	if result[1] < 1 || result[1] > 6 {
		t.Errorf("Expected a number between 1 and 6 but got %d", result[1])
	}
}
