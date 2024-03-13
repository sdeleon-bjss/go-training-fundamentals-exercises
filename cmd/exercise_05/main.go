package main

import "bjss-go-training/pkg"

// Exercise 5:
// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
// - Three single digit numbers from one method.
// - Three double digit numbers from a second method.
// - Three triple digit numbers from a third method.
// - Finally sum all methods into a final sum in the main program.
func main() {
	nums := []int{1, 2, 3, 40, 50, 60, 700, 800, 900}
	numbers := pkg.NumArray{
		Numbers: nums,
	}

	sum := numbers.FinalSum()

	println("Final sum:", sum)
}
