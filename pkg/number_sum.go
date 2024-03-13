package pkg

type NumArray struct {
	Numbers []int
}

func (n *NumArray) SingleDigits() []int {
	return n.Numbers[:3]
}

func (n *NumArray) DoubleDigits() []int {
	return n.Numbers[3:6]
}

func (n *NumArray) TripleDigits() []int {
	return n.Numbers[6:]
}

func (n *NumArray) Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func (n *NumArray) FinalSum() int {
	singleDigits := n.SingleDigits()
	doubleDigits := n.DoubleDigits()
	tripleDigits := n.TripleDigits()

	sum := n.Sum(singleDigits) + n.Sum(doubleDigits) + n.Sum(tripleDigits)

	return sum
}
