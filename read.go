package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the name struct with two fields: fname and lname
type name struct {
	fname string // First name, limited to 20 characters
	lname string // Last name, limited to 20 characters
}

func main() {
	// Prompt user for the name of the text file
	fmt.Print("Enter the name of the text file: ")
	reader := bufio.NewReader(os.Stdin)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the name structs
	var names []name

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		// Check if the line has at least two parts (first name and last name)
		if len(parts) == 2 {
			// Create a new name struct and limit the size of strings to 20 characters
			n := name{
				fname: parts[0],
				lname: parts[1],
			}

			if len(n.fname) > 20 {
				n.fname = n.fname[:20]
			}
			if len(n.lname) > 20 {
				n.lname = n.lname[:20]
			}

			// Add the struct to the slice
			names = append(names, n)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Iterate through the slice and print the first and last names
	fmt.Println("\nNames read from the file:")
	for i, n := range names {
		fmt.Printf("%d. First Name: %s, Last Name: %s\n", i+1, n.fname, n.lname)
	}
}
