//go:build OMIT

package main

import (
	"fmt"
	"sync"
)

// main START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Ping")
		wg.Done()
	}()

	go func() {
		fmt.Println("Pong")
		wg.Done()
	}()

	wg.Wait()
}

// main END OMIT
