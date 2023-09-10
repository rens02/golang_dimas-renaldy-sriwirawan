package main

import (
	"fmt"
)

func SimpleQuations(a, b, c int) {
	check := false
	for x := 1; x < a && !check; x++ {
		// iterate over all possible values of x, from 1 to a-1
		for y := x + 1; y < a && !check; y++ {
			// compute z based on x and y
			z := a - x - y
			// check if the current x, y, and z satisfy the given equations
			if x*y*z == b && x*x+y*y+z*z == c {
				check = true
				fmt.Printf("%d, %d, %d\n", x, y, z)
			}
		}
	}
	if !check {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleQuations(1, 2, 3) // no solution
	SimpleQuations(6, 6, 14)
}
