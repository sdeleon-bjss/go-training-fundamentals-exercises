package main_test

import (
	pkg "bjss-go-training/assessments/fundamentals/04_array_operations"
	"reflect"
	"testing"
)

type testCase struct {
	input    []int
	expected []int
}

func TestAscending(t *testing.T) {
	testCases := []testCase{
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
	testCases := []testCase{
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
