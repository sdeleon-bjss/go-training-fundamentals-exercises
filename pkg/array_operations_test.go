package pkg_test

import (
	"bjss-go-training/pkg"
	"reflect"
	"testing"
)

type numberSorts struct {
	input    []int
	expected []int
}

func TestAscending(t *testing.T) {
	testCases := []numberSorts{
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			input:    []int{2, 0, 2, 4},
			expected: []int{0, 2, 2, 4},
		},
	}

	for _, tc := range testCases {
		result := pkg.Ascending(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func TestDescending(t *testing.T) {
	testCases := []numberSorts{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			input:    []int{2, 0, 2, 4},
			expected: []int{4, 2, 2, 0},
		},
	}

	for _, tc := range testCases {
		result := pkg.Descending(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

type testOddEvens struct {
	scenario      string
	numbers       []int
	direction     string
	expectedOdds  int
	expectedEvens int
}

func TestOddEvenCount(t *testing.T) {
	testCases := []testOddEvens{
		{
			scenario:      "test ascending",
			numbers:       []int{5, 9, 1, 4, 7, 6, 3, 8, 2, 10},
			direction:     "ascending",
			expectedOdds:  25,
			expectedEvens: 30,
		},
		{
			scenario:      "test descending",
			numbers:       []int{5, 9, 1, 4, 7, 6, 3, 8, 2, 10},
			direction:     "descending",
			expectedOdds:  -25,
			expectedEvens: -30,
		},
	}

	for _, tc := range testCases {
		odds, evens := pkg.OddEvenCount(tc.numbers, tc.direction)
		if odds != tc.expectedOdds || evens != tc.expectedEvens {
			t.Errorf("Expected %d odds and %d evens but got %d odds and %d evens", tc.expectedOdds, tc.expectedEvens, odds, evens)
		}
	}
}
