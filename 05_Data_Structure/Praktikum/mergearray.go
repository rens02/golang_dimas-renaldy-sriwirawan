package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := append(arrayA, arrayB...) // combine the two arrays
	keyElement := make(map[string]bool)      // make an empty map to get the key
	result := []string{}

	// create looping to check for duplicates
	for _, val := range mergedArray {
		// make conditions when the keyElement is stiil empty
		if _, ok := keyElement[val]; !ok {
			result = append(result, val)
			keyElement[val] = true
		}
	}

	return result
}

func main() {
	// Test Cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"addie", "steve", "gess"}))
	// [king","devil jin", "akuma","addie","steve","gess" ]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// [sergei , jin , stave , bryan]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshitisu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// [alisa, yoshimitsu, devil jin , law]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// [devil jin, sergei]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// [hworang]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
