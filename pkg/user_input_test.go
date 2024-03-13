package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

func TestGetFullName(t *testing.T) {
	want := "Steven William DeLeon"

	got := pkg.GetFullName("Steven", "William", "DeLeon")

	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}
}
