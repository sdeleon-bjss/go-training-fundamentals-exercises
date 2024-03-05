package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

func TestGetFullName(t *testing.T) {
	first := "Steven"
	middle := "William"
	last := "DeLeon"

	expected := "Steven William DeLeon"

	result := pkg.GetFullName(first, middle, last)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
