package main

import (
	"fmt"
)

func main() {
	// deklarasi variable dan tipe data
	var nilai int64
	// input processing
	fmt.Print("Input nilai: ")
	fmt.Scanln(&nilai)
	
	switch {
	case nilai < 0 || nilai > 100:    // case kurang dari 0 dan melebihi 100
		fmt.Println("Nilai Invalid")
	case nilai >= 80 && nilai <= 100: //case grade A
		fmt.Println("Grade : A")
	case nilai >= 65 && nilai <= 79: //case grade B
		fmt.Println("Grade : B")
	case nilai >= 50 && nilai <= 64: // casse grade C
		fmt.Println("Grade : C")
	case nilai >= 35 && nilai <= 49: //case grade D
		fmt.Println("Grade : D")
	default:
		fmt.Println("Grade : E") // selain semua case diatas
	}

}