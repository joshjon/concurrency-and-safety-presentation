//go:build OMIT

package main

import (
	"fmt"
)

// main START OMIT
func main() {
	go func() { fmt.Println("Ping") }()
	go func() { fmt.Println("Pong") }()
}

// main END OMIT
