//go:build OMIT

package main

import (
	"fmt"
)

// main START OMIT
func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Ping")
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("Pong")
		done <- struct{}{}
	}()

	<-done
	<-done
}

// main END OMIT
