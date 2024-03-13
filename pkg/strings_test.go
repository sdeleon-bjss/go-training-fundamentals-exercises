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
