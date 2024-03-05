package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
	"time"
)

// TODO figure out how to mock and set system time (otherwise this test will fail in future years...)
// (Similar to in jest.setSystemTime and mock)
func TestCalculateAge(t *testing.T) {
	dob := time.Date(1992, time.January, 28, 9, 15, 0, 0, time.UTC)
	result := pkg.CalculateAge(dob)

	expected := 32

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
