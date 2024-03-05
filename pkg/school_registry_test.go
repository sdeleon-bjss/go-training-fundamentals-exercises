package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testPupil struct {
	scenario         string
	pupil            pkg.Pupil
	expectedFullName string
}

func TestNewPupil(t *testing.T) {
	tests := []testPupil{
		{
			scenario:         "Test NewPupil 1",
			pupil:            pkg.NewPupil("Steven", "William", "DeLeon", "01/28/1992", 32),
			expectedFullName: "Steven William DeLeon",
		},
		{
			scenario:         "Test NewPupil 2",
			pupil:            pkg.NewPupil("Gracie", "my", "cat", "06/22/2018", 6),
			expectedFullName: "Gracie my cat",
		},
		{
			scenario:         "Test NewPupil 3",
			pupil:            pkg.NewPupil("Penelope", "my", "cat", "05/12/2017", 7),
			expectedFullName: "Penelope my cat",
		},
		{
			scenario:         "Test NewPupil 4",
			pupil:            pkg.NewPupil("Mr. Boot", "my", "cat", "10/18/2019", 5),
			expectedFullName: "Mr. Boot my cat",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			if test.pupil.GetFullName() != test.expectedFullName {
				t.Errorf("got %s, want %s", test.pupil.GetFullName(), test.expectedFullName)
			}
		})
	}
}
