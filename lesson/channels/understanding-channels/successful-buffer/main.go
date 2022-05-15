package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 12

	fmt.Println(<-c)
}
