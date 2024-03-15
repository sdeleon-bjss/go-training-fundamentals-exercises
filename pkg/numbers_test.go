package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"reflect"
	"testing"
)

var numbers = [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}

func TestIsNumberBetweenRange(t *testing.T) {
	testCases := []struct {
		scenario string
		num      int
		start    int
		end      int
		want     bool
	}{
		{
			scenario: "(success) Number is between the range of 1 and 10",
			num:      5,
			start:    1,
			end:      10,
			want:     true,
		},
		{
			scenario: "(fail) Number is not between range of 1 and 10",
			num:      15,
			start:    1,
			end:      10,
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			got := pkg.IsNumberBetweenRange(tc.num, tc.start, tc.end)

			if got != tc.want {
				t.Errorf("For scenario '%s': Expected %t but got %t", tc.scenario, tc.want, got)
			}
		})
	}
}

func TestNumArray_Sum(t *testing.T) {
	testCase := struct {
		name     string
		numArray *pkg.NumArray
		want     int
	}{
		name: "Should correctly sum numbers using Sum method",
		numArray: &pkg.NumArray{
			Numbers: numbers,
		},
		want: 6,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got := testCase.numArray.Sum(testCase.numArray.SingleDigits())

		if got != testCase.want {
			t.Errorf("Sum() = %d; want %d", got, testCase.want)
		}

	})
}

func TestNumArray_SingleDigits(t *testing.T) {
	testCase := struct {
		name     string
		numArray *pkg.NumArray
		want     []int
	}{
		name: "Should correctly return single digits",
		numArray: &pkg.NumArray{
			Numbers: numbers,
		},
		want: []int{1, 2, 3},
	}

	t.Run(testCase.name, func(t *testing.T) {
		got := testCase.numArray.SingleDigits()

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("SingleDigits() = %v; want %v", got, testCase.want)
		}
	})
}

func TestNumArray_DoubleDigits(t *testing.T) {
	testCase := struct {
		name     string
		numArray *pkg.NumArray
		want     []int
	}{
		name: "Should correctly return double digits",
		numArray: &pkg.NumArray{
			Numbers: numbers,
		},
		want: []int{40, 50, 60},
	}

	t.Run(testCase.name, func(t *testing.T) {
		got := testCase.numArray.DoubleDigits()

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("DoubleDigits() = %v; want %v", got, testCase.want)
		}
	})
}

func TestNumArray_TripleDigits(t *testing.T) {
	testCase := struct {
		name     string
		numArray *pkg.NumArray
		want     []int
	}{
		name: "Should correctly return triple digits",
		numArray: &pkg.NumArray{
			Numbers: numbers,
		},
		want: []int{700, 800, 900},
	}

	t.Run(testCase.name, func(t *testing.T) {
		got := testCase.numArray.TripleDigits()

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("TripleDigits() = %v; want %v", got, testCase.want)
		}
	})
}

func TestNumArray_FinalSum(t *testing.T) {
	testCase := struct {
		name     string
		numArray *pkg.NumArray
		want     int
	}{
		name: "Should correctly return final sum",
		numArray: &pkg.NumArray{
			Numbers: numbers,
		},
		want: 2556,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got := testCase.numArray.FinalSum()

		if got != testCase.want {
			t.Errorf("FinalSum() = %d; want %d", got, testCase.want)
		}
	})
}

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
