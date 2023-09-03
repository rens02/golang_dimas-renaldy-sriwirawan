package main

import (
	"fmt"
)

func PairSum(arr []int, target int) []int {

	mark := make(map[int]int)

	for i, num := range arr {
		// create a variable difference to calculate this target and number
		difference := target - num

		// Making conditions to check the difference is already on the map
		if _, ok := mark[difference]; ok {
			// Return the index of the number of numbers found
			return []int{mark[difference], i}
		}
		// If the difference is not in the map, then save the current numbers and the index
		mark[num] = i
	}

	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))

}
