package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 12
	}()

	fmt.Println(<-c)
}
