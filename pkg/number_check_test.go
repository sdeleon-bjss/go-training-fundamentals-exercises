package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testRangeScenario struct {
	scenario string
	num      int
	start    int
	end      int
	expected bool
}

func TestIsNumberBetweenRange(t *testing.T) {
	testCases := []testRangeScenario{
		{
			scenario: "Number is between the range of 1 and 10",
			num:      5,
			start:    1,
			end:      10,
			expected: true,
		},
		{
			scenario: "Number is not between range of 1 and 10",
			num:      15,
			start:    1,
			end:      10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			result := pkg.IsNumberBetweenRange(tc.num, tc.start, tc.end)
			if result != tc.expected {
				t.Errorf("For scenario '%s': Expected %t but got %t", tc.scenario, tc.expected, result)
			}
		})
	}
}
