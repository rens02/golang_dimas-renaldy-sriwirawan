package main

import (
	"fmt"
)

func getMinMax(numbers ...int) (min int, max int) {
	// Initialize minimum and maximum values
	min = numbers[0]
	max = numbers[0]

	// Iterate through slice numbers using range
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6 int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max := getMinMax(a1, a2, a3, a4, a5, a6)

	// find the average, minimum, and maximum values
	minPtr := &min
	maxPtr := &max
	fmt.Println("Nilai Min", *minPtr)
	fmt.Println("Nilai Max", *maxPtr)

}
