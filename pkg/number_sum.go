package pkg

func CalculateSum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
