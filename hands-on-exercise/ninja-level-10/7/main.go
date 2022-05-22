package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go genNumbers(c)

	receive(c)
}

func genNumbers(c chan int) {
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			defer wg.Done()
			for i := 1; i <= 10; i++ {
				c <- rand.Intn(100)
			}
		}()
	}
	wg.Wait()
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
