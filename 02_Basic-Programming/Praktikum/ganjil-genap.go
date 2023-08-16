package main

import (
	"fmt"
)

func main() {
	// deklarasi variabel dan type data
	var number int64
	// meminta input angka
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&number)
	
	//  jika angka dapat dibagi 2 maka angka tersebut adalah genap
	if number%2 == 0 {
		fmt.Println(number, " adalah genap")
	} else {
		fmt.Println(number, " adalah ganjil")
	}

}
