//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	counter := int32(0)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// main END OMIT
