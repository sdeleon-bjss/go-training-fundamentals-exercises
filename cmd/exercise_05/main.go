package main

import "bjss-go-training/pkg"

// Exercise 5:
// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
// - Three single digit numbers from one method.
// - Three double digit numbers from a second method.
// - Three triple digit numbers from a third method.
// - Finally sum all methods into a final sum in the main program.
func main() {
	numbers := []int{1, 2, 3, 10, 20, 30, 100, 200, 300}
	singleDigits := numbers[:3]
	doubleDigits := numbers[3:6]
	tripleDigits := numbers[6:]

	singleDigitSum := pkg.CalculateSum(singleDigits)
	doubleDigitSum := pkg.CalculateSum(doubleDigits)
	tripleDigitSum := pkg.CalculateSum(tripleDigits)

	finalDigits := []int{singleDigitSum, doubleDigitSum, tripleDigitSum}
	finalSum := pkg.CalculateSum(finalDigits)

	println(finalSum)

}
