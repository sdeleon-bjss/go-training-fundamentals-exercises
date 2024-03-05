package main_test

import (
	pkg "bjss-go-training/assessments/fundamentals/06_age_calculator"
	"testing"
	"time"
)

// TODO figure out how to mock and set system time (otherwise this test will fail in future years...)
func TestCalculateAge(t *testing.T) {
	dob := time.Date(1992, time.January, 28, 9, 15, 0, 0, time.UTC)
	result := pkg.CalculateAge(dob)

	expected := 32

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
