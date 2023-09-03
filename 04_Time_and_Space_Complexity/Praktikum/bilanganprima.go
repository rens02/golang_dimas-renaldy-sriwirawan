package main

import (
	"fmt"
)

func primeNumber(number int) bool {

	//  check if the number is greater than 1
	if number <= 1 {
		return false
	}
	// check if the number is less than 3
	if number <= 3 {
		return true
	}
	// Check if the number is divisible by 2 / 3
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	//  the process of finding the quadrant root of the number
	for i := 5; i*i <= number; i += 6 {
		//  check if the number is divisible by i / or i % 2
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
