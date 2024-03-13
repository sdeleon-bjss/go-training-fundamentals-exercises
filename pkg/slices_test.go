package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
)

func TestName_SetFullName(t *testing.T) {
	testCases := []struct {
		name          pkg.Name
		fullNameInput string
		want          bool
	}{
		{name: pkg.Name{}, fullNameInput: "Steven William DeLeon", want: false},
		{name: pkg.Name{}, fullNameInput: "Steven William", want: true},
		{name: pkg.Name{}, fullNameInput: "foo bar baz", want: false},
	}

	for _, tc := range testCases {
		err := tc.name.SetFullName(tc.fullNameInput)

		if (err != nil) != tc.want {
			t.Errorf("Name.SetFullName(%s) = %v, want %v", tc.fullNameInput, err, tc.want)
		}
	}
}

func TestName_GetFullName(t *testing.T) {
	name := pkg.Name{
		First:  "Steven",
		Middle: "William",
		Last:   "DeLeon",
	}

	got := name.GetFullName()
	want := "Steven William DeLeon"

	if got != want {
		t.Errorf("Name.GetFullName() = %s, want %s", got, want)
	}
}
