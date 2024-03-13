package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

func TestNewPupil(t *testing.T) {
	testCases := []struct {
		firstName     string
		middleName    string
		lastName      string
		dateOfBirth   string
		expectedError bool
	}{
		{"Archie", "James", "Johnson", "04/15/2013", false},
		{"", "James", "Johnson", "04/15/2013", true},
		{"Archie", "", "Johnson", "04/15/2013", true},
		{"Archie", "James", "", "04/15/2013", true},
		{"Archie", "James", "Johnson", "04/15/20132", true},
	}

	for _, tc := range testCases {
		_, err := pkg.NewPupil(tc.firstName, tc.middleName, tc.lastName, tc.dateOfBirth)
		if (err != nil) != tc.expectedError {
			t.Errorf("Expected error: %v but got %v", tc.expectedError, err)
		}
	}
}

func TestPupil_GetFullName(t *testing.T) {
	pupil, _ := pkg.NewPupil("Archie", "James", "Johnson", "04/15/2013")

	if pupil.GetFullName() != "Archie James Johnson" {
		t.Errorf("Expected 'Archie James Johnson' but got %s", pupil.GetFullName())
	}
}

func TestPupil_GetDateOfBirth(t *testing.T) {
	pupil, _ := pkg.NewPupil("Archie", "James", "Johnson", "04/15/2013")

	if pupil.GetDateOfBirth() != "04/15/2013" {
		t.Errorf("Expected '04/15/2013' but got %s", pupil.GetDateOfBirth())
	}
}

func TestPupil_GetAge(t *testing.T) {
	pupil, _ := pkg.NewPupil("Archie", "James", "Johnson", "04/15/2013")

	if pupil.GetAge() != 10 {
		t.Errorf("Expected 11 but got %d", pupil.GetAge())
	}
}
