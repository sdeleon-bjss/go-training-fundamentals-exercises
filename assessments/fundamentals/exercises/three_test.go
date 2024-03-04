package exercises_test

import (
	"bjss-go-training/assessments/fundamentals/exercises"
	"testing"
)

func TestIsNumberBetweenRange(t *testing.T) {
	num := 5
	rangeStart := 1
	rangeEnd := 10

	expected := true

	result := exercises.IsNumberBetweenRange(num, rangeStart, rangeEnd)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	num = 0
	rangeStart = 1
	rangeEnd = 10

	expected = false

	result = exercises.IsNumberBetweenRange(num, rangeStart, rangeEnd)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}
