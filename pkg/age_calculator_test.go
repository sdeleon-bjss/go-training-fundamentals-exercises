package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
	"time"
)

var mockedTime = time.Date(2024, time.March, 12, 0, 0, 0, 0, time.UTC)

type testCandidate struct {
	dob        time.Time
	timeFrozen time.Time
	want       int
}

func TestPerson_Age(t *testing.T) {
	testCases := []testCandidate{
		{
			dob:        time.Date(1992, time.January, 28, 0, 0, 0, 0, time.UTC),
			timeFrozen: mockedTime,
			want:       32,
		},
		{
			dob:        time.Date(1992, time.August, 31, 0, 0, 0, 0, time.UTC),
			timeFrozen: mockedTime,
			want:       31,
		},
		{
			dob:        time.Date(1975, time.May, 21, 0, 0, 0, 0, time.UTC),
			timeFrozen: mockedTime,
			want:       48,
		},
		{
			dob:        time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
			timeFrozen: mockedTime,
			want:       3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.dob.String(), func(t *testing.T) {
			p := &pkg.Person{DOB: tc.dob}
			got := p.Age()

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
