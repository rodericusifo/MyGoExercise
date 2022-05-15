package main

import "fmt"

func main() {
	c := make(chan int)
	d := genNumbers(c)

	receive(d)
}

func genNumbers(c chan int) <-chan int {
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}