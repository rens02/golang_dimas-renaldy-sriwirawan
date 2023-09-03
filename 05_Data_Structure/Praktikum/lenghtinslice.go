package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	// create a result variable with a map type to accomodate a unique mapping
	result := make(map[string]int)
	// Create a loop to fill the key with the value input
	for _, s := range slice {
		// Add the key and if the key is already there, do the increment for the value
		result[s]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))

}
