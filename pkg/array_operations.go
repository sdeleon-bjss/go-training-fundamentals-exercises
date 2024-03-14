package pkg

import (
	"sort"
	"strings"
)

// Ascending sorts the numbers in ascending order
func Ascending(numbers []int) []int {
	ascending := make([]int, len(numbers))
	copy(ascending, numbers)

	sort.Ints(ascending)
	return ascending
}

// Descending sorts the numbers in descending order
func Descending(numbers []int) []int {
	descending := make([]int, len(numbers))
	copy(descending, numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(descending)))
	return descending
}

// OddEvenCount counts the odd and even numbers in the array, and returns the sum of each
//
// The direction parameter can be either "ascending" or "descending"
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
