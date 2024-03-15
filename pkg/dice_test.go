package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
)

func TestDice_Roll(t *testing.T) {
	dice := pkg.Dice{}

	for i := 0; i < 50; i++ {
		got := dice.Roll()
		if !pkg.IsNumberBetweenRange(got, 1, 6) {
			t.Errorf("expected got to be between 1 and 6, got %d", got)
		}
	}
}

func TestDice_SetCombined(t *testing.T) {
	dice := pkg.Dice{FirstRoll: 3, SecondRoll: 4}
	got := dice.SetCombined()

	if got != 7 {
		t.Errorf("expected got to be 7, got %d", got)
	}
}

func TestDice_ProcessRoll(t *testing.T) {
	testCases := []struct {
		name string
		want string
		dice pkg.Dice
	}{
		{name: "Natural seven", want: "NATURAL", dice: pkg.Dice{FirstRoll: 3, SecondRoll: 4, Combined: 7}},
		{name: "Natural eleven", want: "NATURAL", dice: pkg.Dice{FirstRoll: 5, SecondRoll: 6, Combined: 11}},
		{name: "Snake Eyes", want: "SNAKE-EYES-CRAPS", dice: pkg.Dice{FirstRoll: 1, SecondRoll: 1, Combined: 2}},
		{name: "Loss three", want: "LOSS-CRAPS", dice: pkg.Dice{FirstRoll: 1, SecondRoll: 2, Combined: 3}},
		{name: "Loss twelve", want: "LOSS-CRAPS", dice: pkg.Dice{FirstRoll: 6, SecondRoll: 6, Combined: 12}},
		{name: "Neutral", want: "NEUTRAL", dice: pkg.Dice{FirstRoll: 2, SecondRoll: 3, Combined: 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.dice.ProcessRoll()

			if got != tc.want {
				t.Errorf("expected got to be %s, got %s", tc.want, got)
			}
		})
	}

}
