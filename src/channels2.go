//go:build OMIT

package main

import (
	"fmt"
)

// main START OMIT

func main() {
	done := make(chan struct{})
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch2 <- "Ping"
		fmt.Println(<-ch1)
		done <- struct{}{}
	}()

	go func() {
		fmt.Println(<-ch2)
		ch1 <- "Pong"
		done <- struct{}{}
	}()

	<-done
	<-done
}

// main END OMIT
