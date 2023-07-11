//go:build OMIT

package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		safeCounter()
	}
}

// main START OMIT
// Execute 10 times

func safeCounter() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// main END OMIT
