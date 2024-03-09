package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testPupil struct {
	scenario            string
	pupil               []pkg.Person
	expectedFullName    []string
	expectedDateOfBirth []string
	expectedAge         []int
}

func TestNewPupil(t *testing.T) {
	tests := []testPupil{
		{
			scenario: "Create a new pupil",
			pupil: []pkg.Person{
				pkg.NewPupil("Steven", "William", "DeLeon", "01/28/1992", 32),
				pkg.NewPupil("Gracie", "my", "cat", "06/22/2018", 6),
				pkg.NewPupil("Penelope", "my", "cat", "05/12/2017", 7),
				pkg.NewPupil("Mr. Boot", "my", "cat", "10/18/2019", 5),
			},
			expectedFullName: []string{
				"Steven William DeLeon",
				"Gracie my cat",
				"Penelope my cat",
				"Mr. Boot my cat",
			},
			expectedDateOfBirth: []string{
				"01/28/1992",
				"06/22/2018",
				"05/12/2017",
				"10/18/2019",
			},
			expectedAge: []int{
				32,
				6,
				7,
				5,
			},
		},
	}

	for _, test := range tests {
		for i, pupil := range test.pupil {
			if pupil.GetFullName() != test.expectedFullName[i] {
				t.Errorf("Test: %s - Expected: %s, Got: %s", test.scenario, test.expectedFullName[i], pupil.GetFullName())
			}
			if pupil.GetDateOfBirth() != test.expectedDateOfBirth[i] {
				t.Errorf("Test: %s - Expected: %s, Got: %s", test.scenario, test.expectedDateOfBirth[i], pupil.GetDateOfBirth())
			}
			if pupil.GetAge() != test.expectedAge[i] {
				t.Errorf("Test: %s - Expected: %d, Got: %d", test.scenario, test.expectedAge[i], pupil.GetAge())
			}
		}
	}
}
