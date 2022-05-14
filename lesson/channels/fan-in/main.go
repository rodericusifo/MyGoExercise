package main

import (
	"fmt"
	"sync"
)

// NOTE: Fan-In is Taking values from many channels, and putting those values onto one channel.

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
