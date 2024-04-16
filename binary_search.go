package main

import (
	"fmt"
	"sort"
)

func binarySearch(array []int, key int) int {

	//Set low and high boundaries
	lowerBound := 0
	upperBound := len(array) - 1

	//We begin a loop in which we keep inspecting the middlemost value between the upper and lower bound
	for lowerBound <= upperBound {

		//Find midpoint between lower and upper bound
		midPoint := (lowerBound + upperBound) / 2

		// Inspect value at Midpoint
		valueAtMidpoint := array[midPoint]

		if key == valueAtMidpoint {
			return midPoint
		} else if key < valueAtMidpoint {
			// if target is lower, focus on the lesser half
			upperBound = midPoint - 1
		} else {
			//if target is greater focus on the higher half
			lowerBound = midPoint + 1
		}

	}

	// If element not found, return -1
	return -1

}

func main() {
	fmt.Println(">>> Binary Search in an Array <<<")

	// Array of numbers
	numbers := []int{6, 9, 2, 8, 3, 5, 4}
	key := 1

	//Binary search works on Ordered Arrays, Sort the arrays
	sort.Ints(numbers)

	fmt.Println("Sorted array: ", numbers)

	results := binarySearch(numbers, key)

	if results != -1 {
		fmt.Printf("Element found at index %d.\n\n", results)
	} else {
		fmt.Println("Element not found in array")
	}

}
