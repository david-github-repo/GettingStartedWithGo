package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for name
	fmt.Print("Enter a name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	name = strings.TrimSpace(name)

	// Prompt for address
	fmt.Print("Enter an address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	address = strings.TrimSpace(address)

	// Create a map with name and address
	personMap := map[string]string{
		"name":    name,
		"address": address,
	}

	// Marshal the map into JSON
	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Print the JSON object
	fmt.Println(string(jsonData))
}
