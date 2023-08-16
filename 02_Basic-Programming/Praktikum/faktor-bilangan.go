package main

import (
	"fmt"
)

func main() {

    var number int
    fmt.Print("angka: ")
    fmt.Scanln(&number)


    for i := 1; i <= number; i++ {
        if number%i == 0 {
            fmt.Printf("%d ", i)
            fmt.Println()
        }
       
    }
}