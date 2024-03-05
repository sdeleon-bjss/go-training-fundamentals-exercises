package pkg

import "sort"

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
