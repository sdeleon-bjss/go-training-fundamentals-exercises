package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"reflect"
	"testing"
)

func TestAscending(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			input: []int{2, 0, 2, 4},
			want:  []int{0, 2, 2, 4},
		},
	}

	for _, tc := range testCases {
		got := pkg.Ascending(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}

func TestDescending(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			input: []int{2, 0, 2, 4},
			want:  []int{4, 2, 2, 0},
		},
	}

	for _, tc := range testCases {
		got := pkg.Descending(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}

type testOddEvens struct {
	scenario      string
	inputNumbers  []int
	sortDirection string
	wantOdds      int
	wantEvens     int
}

func TestOddEvenCount(t *testing.T) {
	testCases := []testOddEvens{
		{
			scenario:      "test ascending",
			inputNumbers:  []int{5, 9, 1, 4, 7, 6, 3, 8, 2, 10},
			sortDirection: "ascending",
			wantOdds:      25,
			wantEvens:     30,
		},
		{
			scenario:      "test descending",
			inputNumbers:  []int{5, 9, 1, 4, 7, 6, 3, 8, 2, 10},
			sortDirection: "descending",
			wantOdds:      -25,
			wantEvens:     -30,
		},
	}

	for _, tc := range testCases {
		gotOdds, gotEvens := pkg.OddEvenCount(tc.inputNumbers, tc.sortDirection)
		if gotOdds != tc.wantOdds || gotEvens != tc.wantEvens {
			t.Errorf("Expected %d odds and %d evens but got %d odds and %d evens", tc.wantOdds, tc.wantEvens, gotOdds, gotEvens)
		}
	}
}
