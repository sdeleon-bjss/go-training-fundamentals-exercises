package exercises

import "testing"

func TestCalculateSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 15
	result := calculateSum(nums)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	nums = []int{10, 20, 30, 40, 50}
	expected = 150

	result = calculateSum(nums)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestNineNumberSum(t *testing.T) {
	singleDigits := []int{1, 1, 1}
	doubleDigits := []int{10, 10, 10}
	tripleDigits := []int{100, 100, 100}

	expected := 333
	result := nineNumberSum(singleDigits, doubleDigits, tripleDigits)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
