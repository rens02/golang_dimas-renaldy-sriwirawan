package main

import (
	"fmt"
)

func main() {

	var number int

 	fmt.Print("Input number : ")
	fmt.Scanln(&number)
	fmt.Println("Output :")
	for i := 0; i <= number; i++ {
		for j := 1; j <= number-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
