package exercises

// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
//Three single digit numbers from one method.
//Three double digit numbers from a second method.
//Three triple digit numbers from a third method.
//Finally sum all methods into a final sum in the main program.

func nineNumberSum(singles, doubles, triples []int) int {
	singlesSum := calculateSum(singles)
	doublesSum := calculateSum(doubles)
	triplesSum := calculateSum(triples)

	return singlesSum + doublesSum + triplesSum
}

func calculateSum(nums []int) int {
	sum := 0

	for _, i := range nums {
		sum += i
	}

	return sum
}
