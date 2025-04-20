package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Create an integer slice with capacity 3
	intSlice := make([]int, 0, 3)

	for {
		// Prompt user for input
		fmt.Print("Enter an integer to add to the slice or 'X' to quit: ")
		var input string
		fmt.Scan(&input)

		// Check if user wants to quit
		if input == "X" || input == "x" {
			break
		}

		// Convert input to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X' to quit.")
			continue
		}

		// Add integer to slice
		intSlice = append(intSlice, num)

		// Sort the slice
		sort.Ints(intSlice)

		// Print the sorted slice
		fmt.Println("Sorted slice:", intSlice)
	}
}
