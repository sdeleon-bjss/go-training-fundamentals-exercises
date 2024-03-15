package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"testing"
)

func TestConcatStrings(t *testing.T) {
	s1 := "Hello"
	s2 := " "
	s3 := "World"

	expected := "Hello World"

	result := pkg.ConcatStrings(s1, s2, s3)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestGetFullName(t *testing.T) {
	want := "Steven William DeLeon"

	got := pkg.GetFullName("Steven", "William", "DeLeon")

	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}
}

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
