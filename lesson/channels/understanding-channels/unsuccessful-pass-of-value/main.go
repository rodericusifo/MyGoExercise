package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 12

	fmt.Println(<-c)
}
