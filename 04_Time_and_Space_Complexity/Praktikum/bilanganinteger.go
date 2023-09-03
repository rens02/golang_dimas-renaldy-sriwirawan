package main

import (
	"fmt"
)

func pow(x, n int) int {
	// hendling check n == 0 if true returns 1
	if n == 0 {
		return 1
	}
	// decalration variable result = 1
	result := 1
	for n > 0 {
		// check if the exponent of in is odd
		if n%2 == 1 {
			result *= x // multiply the result by the natural number if the exponent is odd
		}

		x *= x // square the natural number at each iteration
		n /= 2 // Decrease the power by diving by 2 in each iteration
	}
	return result

}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
