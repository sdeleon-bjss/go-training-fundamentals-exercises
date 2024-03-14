package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
)

func TestIsNumberBetweenRange(t *testing.T) {
	testCases := []struct {
		scenario string
		num      int
		start    int
		end      int
		want     bool
	}{
		{
			scenario: "(success) Number is between the range of 1 and 10",
			num:      5,
			start:    1,
			end:      10,
			want:     true,
		},
		{
			scenario: "(fail) Number is not between range of 1 and 10",
			num:      15,
			start:    1,
			end:      10,
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			got := pkg.IsNumberBetweenRange(tc.num, tc.start, tc.end)

			if got != tc.want {
				t.Errorf("For scenario '%s': Expected %t but got %t", tc.scenario, tc.want, got)
			}
		})
	}
}
