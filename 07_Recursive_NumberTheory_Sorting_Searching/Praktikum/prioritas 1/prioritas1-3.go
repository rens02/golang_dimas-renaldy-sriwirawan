package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	count := 0
	lastPrime := 0
	for i := 2; count < number; i++ {
		isPrime := true
		limit := int(math.Sqrt(float64(i))) + 1
		for j := 2; j < limit; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			lastPrime = i
			count++
		}
	}
	return lastPrime
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
