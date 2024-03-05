package main_test

import (
	pkg "bjss-go-training/assessments/fundamentals/01_strings"
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
