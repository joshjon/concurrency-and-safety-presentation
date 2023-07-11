//go:build OMIT

package main

import (
	"fmt"
	"sync"
)

// main START OMIT
func main() {
	var value string
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		value = "foo"
	}()

	go func() {
		defer wg.Done()
		value = "bar"
	}()

	wg.Wait()
	fmt.Println(value)
}

// main END OMIT
