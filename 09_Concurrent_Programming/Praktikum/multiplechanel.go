package main

import (
	"fmt"
	"time"
)

func printMultiplesOfThree(ch chan int) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printMultiplesOfThree(ch)

	for num := range ch {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println(num)
	}
}
