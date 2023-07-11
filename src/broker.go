//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		publishCh: make(chan T, 1),
		subCh:     make(chan chan T, 1),
		unsubCh:   make(chan chan T, 1),
	}
}

// Broker START OMIT
type Broker[T any] struct {
	publishCh chan T
	subCh     chan chan T
	unsubCh   chan chan T
}

func (b *Broker[T]) Publish(msg T) {
	b.publishCh <- msg
}

func (b *Broker[T]) Subscribe() chan T {
	msgCh := make(chan T, 5)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker[T]) Unsubscribe(msgCh chan T) {
	b.unsubCh <- msgCh
}

// Broker END OMIT

// Start START OMIT
func (b *Broker[T]) Start() {
	go func() {
		subs := map[chan T]struct{}{} // set of subscriber channels

		for {
			select {
			case msgCh := <-b.subCh:
				subs[msgCh] = struct{}{}
			case msgCh := <-b.unsubCh:
				delete(subs, msgCh)
			case msg := <-b.publishCh:
				for msgCh := range subs {
					msgCh <- msg
				}
			}
		}
	}()
}

// Start END OMIT

// PubSub START OMIT

// PubSub END OMIT

// main START OMIT
func main() {
	b := NewBroker[string]()
	b.Start()

	for i := 0; i < 5; i++ {
		go func(id int) {
			sub := b.Subscribe()
			for msg := range sub {
				fmt.Printf("sub %d received: %s\n", id, msg)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		b.Publish(fmt.Sprintf("message %d", i))
		time.Sleep(time.Second)
	}
}

// main END OMIT
