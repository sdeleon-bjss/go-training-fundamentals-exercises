package pkg

// NumArray struct to storing and working with numbers
//
// Numbers: Array of 9 numbers
type NumArray struct {
	Numbers [9]int
}

// SingleDigits method to return the first three numbers
func (n *NumArray) SingleDigits() []int {
	return n.Numbers[:3]
}

// DoubleDigits method to return the second three numbers
func (n *NumArray) DoubleDigits() []int {
	return n.Numbers[3:6]
}

// TripleDigits method to return the last three numbers
func (n *NumArray) TripleDigits() []int {
	return n.Numbers[6:]
}

// Sum method to sum the numbers
func (n *NumArray) Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// FinalSum method to sum all the numbers from the three methods
func (n *NumArray) FinalSum() int {
	singleDigits := n.SingleDigits()
	doubleDigits := n.DoubleDigits()
	tripleDigits := n.TripleDigits()

	sum := n.Sum(singleDigits) + n.Sum(doubleDigits) + n.Sum(tripleDigits)

	return sum
}
