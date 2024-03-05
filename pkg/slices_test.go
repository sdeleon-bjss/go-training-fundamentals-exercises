package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testName struct {
	scenario string
	name     pkg.Name
	expected string
}

func TestName_GetFullName(t *testing.T) {
	testCases := []testName{
		{
			scenario: "Should successfully generate outcome using my name",
			name: pkg.Name{
				First:  "Steven",
				Middle: "William",
				Last:   "DeLeon",
			},
			expected: "Steven William DeLeon",
		},
		{
			scenario: "Should successfully generate outcome despite length or repetitive chars",
			name: pkg.Name{
				First:  "aaa",
				Middle: "bbbbbbb",
				Last:   "ccccc",
			},
			expected: "aaa bbbbbbb ccccc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			result := tc.name.GetFullName()
			if result != tc.expected {
				t.Errorf("For scenario '%s': Expected %s but got %s", tc.scenario, tc.expected, result)
			}
		})
	}
}
