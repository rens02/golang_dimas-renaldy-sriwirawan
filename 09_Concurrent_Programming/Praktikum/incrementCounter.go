package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementCounter(&wg)
	go incrementCounter(&wg)

	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)

}

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
