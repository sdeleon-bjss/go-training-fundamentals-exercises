package exercises

import (
	"sort"
)

// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// Display the array content in ascending sequential order 1 to 10.
// Display the array content in descending sequential order 10 to 1.

// TODO
// Count even numbers and odd numbers in increasing and decreasing sequential order.
// Display the even and odd count sequences to screen.

type Numbers struct {
	Value []int
}

func (n *Numbers) Init() {
	n.Value = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (n *Numbers) Ascending() []int {
	ascending := make([]int, len(n.Value))
	copy(ascending, n.Value)
	sort.Ints(ascending)

	return ascending
}

func (n *Numbers) Descending() []int {
	descending := make([]int, len(n.Value))
	copy(descending, n.Value)

	// reference: https://golangcookbook.com/chapters/arrays/reverse/
	for i := 0; i < len(descending)/2; i++ {
		j := len(descending) - i - 1
		descending[i], descending[j] = descending[j], descending[i]
	}

	return descending
}
