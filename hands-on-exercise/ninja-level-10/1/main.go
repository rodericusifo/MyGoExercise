package main

import (
	"fmt"
)

func main() {
	// Pass of Value
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Buffer
	d := make(chan int, 1)

	d <- 42

	fmt.Println(<-d)
}
