package main

import (
	"fmt"
)

func binaryArray(n int) []int {

	ans := make([]int, n+1)
	// problem solving with brute force
	for i := 0; i <= n; i++ {
		binary := 0
		pow := 1
		for j := i; j > 0; j /= 2 {
			binary += pow * (j % 2)
			pow *= 10
		}
		ans[i] = binary
	}
	return ans
}

func main() {
	ans := binaryArray(3)
	fmt.Println(ans)
}
