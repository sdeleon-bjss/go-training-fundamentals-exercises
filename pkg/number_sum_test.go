package pkg_test

import (
	"github.com/sdeleon-bjss/pkg"
	"reflect"
	"testing"
)

var numbers = [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}

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
