package pkg

import (
	"sort"
	"strings"
)

func IsNumberBetweenRange(num int, start int, end int) bool {
	if num >= start && num <= end {
		return true
	}

	return false
}

type NumArray struct {
	Numbers [9]int
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

func Ascending(numbers []int) []int {
	ascending := make([]int, len(numbers))
	copy(ascending, numbers)

	sort.Ints(ascending)
	return ascending
}

func Descending(numbers []int) []int {
	descending := make([]int, len(numbers))
	copy(descending, numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(descending)))
	return descending
}

func OddEvenCount(numbers []int, direction string) (int, int) {
	even := 0
	odd := 0
	direction = strings.ToLower(direction)

	if direction == "ascending" {
		for _, num := range numbers {
			if num%2 == 0 {
				even += num
				continue
			}
			odd += num
		}
	}

	if direction == "descending" {
		for i := len(numbers) - 1; i >= 0; i-- {
			if numbers[i]%2 != 0 {
				odd -= numbers[i]
			}

			if numbers[i]%2 == 0 {
				even -= numbers[i]
			}
		}
	}

	return odd, even
}
