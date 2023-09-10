package main

import "fmt"

func getRows(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	// create a triangular variable for numRows-1 row
	triangle := getRows(numRows - 1)

	// Initialize the array on the nth line
	row := make([]int, numRows)
	row[0] = 1
	row[numRows-1] = 1

	// Loop for each element in the nth row
	for i := 1; i < numRows-1; i++ {
		// Sets elements according to Pascal's triangle rule
		row[i] = triangle[numRows-2][i-1] + triangle[numRows-2][i]
	}

	// Merge the 2 dimensional array in row numRows-1 with the array in row n
	triangle = append(triangle, row)

	return triangle
}

func main() {
	result := getRows(5)
	fmt.Println(result)
}
