package exercise_01_strings

import "testing"

func TestConcatStrings(t *testing.T) {
	stringOne := "Hello there"
	stringTwo := "my name is"
	stringThree := "Steven!"

	expected := "Hello there my name is Steven!"

	result := ConcatStrings(stringOne, " ", stringTwo, " ", stringThree)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
