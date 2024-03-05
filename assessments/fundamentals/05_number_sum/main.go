package main

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

	singleDigitSum := CalculateSum(singleDigits)
	doubleDigitSum := CalculateSum(doubleDigits)
	tripleDigitSum := CalculateSum(tripleDigits)

	finalDigits := []int{singleDigitSum, doubleDigitSum, tripleDigitSum}
	finalSum := CalculateSum(finalDigits)

	println(finalSum)

}

func CalculateSum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
