package main

import (
	"fmt"
	"time"
)

func printMultiplex(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	go printMultiplex(2)
	time.Sleep(3 * time.Second)
}
