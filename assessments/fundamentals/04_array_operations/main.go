package main

import (
	"fmt"
	"sort"
)

// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// - Display the array content in ascending sequential order 1 to 10.
// - Display the array content in descending sequential order 10 to 1.
// TODO
// - Count even numbers and odd numbers in increasing and decreasing sequential order.
// - Display the even and odd count sequences to screen.

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ascending := Ascending(numbers)
	descending := Descending(numbers)

	fmt.Println("Ascending:", ascending)
	fmt.Println("Descending:", descending)
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

// TODO part 3 and 4 + tests
