//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// Scheduler START OMIT
type Scheduler struct {
	workers chan struct{} // Acts like a semaphore
}

func (s *Scheduler) Add(item any, job func(item any)) {
	s.workers <- struct{}{} // Spawn a worker
	go func() {
		job(item)
		<-s.workers // Remove worker
	}()
}

// Scheduler END OMIT

// main START OMIT
func main() {
	s := Scheduler{workers: make(chan struct{}, 3)}

	for i := 0; i < 15; i++ {
		s.Add(i, handle)
	}

	wg.Wait()
}

// main END OMIT

var wg sync.WaitGroup

func handle(item any) {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("handled %v\n", item)
}
