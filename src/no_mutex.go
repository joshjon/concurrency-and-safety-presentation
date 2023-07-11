//go:build OMIT

package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		unsafeCounter()
	}
}

// main START OMIT
// Execute 10 times

func unsafeCounter() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// main END OMIT
