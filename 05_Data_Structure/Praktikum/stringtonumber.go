package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	// Create a variable of the thought to accommodate
	// the string numbers sent here I use the key rune maping the unicode character
	lntAngka := make(map[rune]int)

	// Calculate the appearance of each number
	for _, c := range angka {
		lntAngka[c]++
	}

	//  make looping to find unique numbers / not duplicate
	result := []int{}
	for _, c := range angka {
		// make a check to check whether the number is only one / not duplicate
		if lntAngka[c] == 1 {
			// append result variable with Conversion from the Rune Value
			result = append(result, int(c-'0'))
		}
	}

	return result

}

func main() {

	fmt.Println((munculSekali("1234123")))    // [4]
	fmt.Println((munculSekali("76523752")))   // [6 3]
	fmt.Println((munculSekali("12345")))      // [1 2 3 4 5]
	fmt.Println((munculSekali("1122334455"))) // []
	fmt.Println((munculSekali("0872504")))    // [8 7 2 5 4]

}
