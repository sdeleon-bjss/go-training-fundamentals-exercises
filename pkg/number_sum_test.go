package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	numbers := []int{1, 2, 3, 10, 20, 30, 100, 200, 300}
	singleDigits := numbers[:3]
	doubleDigits := numbers[3:6]
	tripleDigits := numbers[6:]

	expectedSingleDigitSum := 6
	expectedDoubleDigitSum := 60
	expectedTripleDigitSum := 600
	expectedFinal := 666

	singleDigitSum := pkg.CalculateSum(singleDigits)
	doubleDigitSum := pkg.CalculateSum(doubleDigits)
	tripleDigitSum := pkg.CalculateSum(tripleDigits)

	finalDigits := []int{singleDigitSum, doubleDigitSum, tripleDigitSum}
	finalSum := pkg.CalculateSum(finalDigits)

	if singleDigitSum != expectedSingleDigitSum {
		t.Errorf("Expected %d but got %d", expectedSingleDigitSum, singleDigitSum)
	}

	if doubleDigitSum != expectedDoubleDigitSum {
		t.Errorf("Expected %d but got %d", expectedDoubleDigitSum, doubleDigitSum)
	}

	if tripleDigitSum != expectedTripleDigitSum {
		t.Errorf("Expected %d but got %d", expectedTripleDigitSum, tripleDigitSum)
	}

	if finalSum != expectedFinal {
		t.Errorf("Expected %d but got %d", expectedFinal, finalSum)
	}
}
