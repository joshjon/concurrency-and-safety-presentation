//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool START OMIT
type WorkerPool struct{ jobs chan func() }

func (p *WorkerPool) Start(count int) {
	for i := 0; i < count; i++ {
		go func() {
			for j := range p.jobs {
				j()
			}
		}()
	}
}

// WorkerPool END OMIT

// main START OMIT
func main() {
	p := WorkerPool{jobs: make(chan func(), 10)}
	p.Start(3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		p.jobs <- job
	}

	wg.Wait()
}

// main END OMIT

var wg sync.WaitGroup

func job() {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("job finished")
}
