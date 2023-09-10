package main

import (
	"fmt"
	"time"
)

func printMultiplesOfThree(ch chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go printMultiplesOfThree(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
		time.Sleep(3000 * time.Millisecond)
	}
	close(ch)
}
