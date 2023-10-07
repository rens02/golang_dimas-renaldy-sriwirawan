package main

import (
	"fmt"
)

func maxSequence(arr []int) int {
	var sumMax, sumRange int= 0,0
	
	for _, num := range arr {
		sumRange += num
		if sumRange < 0 {
			sumRange = 0
		}
		if sumMax < sumRange {
			sumMax = sumRange
		}
	}
	
	return sumMax
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) 
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) 
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) 
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) 
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) 
}
