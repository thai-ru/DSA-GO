package main

import "fmt"

func linearSearch(arr []int, key int) int { // Pas the array and target key

	// Running a for loop over the array
	for i := 0; i < len(arr); i++ {

		//If key is found
		if arr[i] == key {
			return i
		}
	}

	// else return -1 if not found
	return -1
}

func main() {

	fmt.Println("<<<< Linear Search >>>>>")

	// Array of numbers
	arr := []int{6, 9, 2, 8, 3, 5, 4}

	// Change value of target
	key := 2

	result := linearSearch(arr, key)

	if result != -1 {
		fmt.Printf("Target found at index %d \n", result)
	} else {
		fmt.Println("Target not found in array")
	}

}
