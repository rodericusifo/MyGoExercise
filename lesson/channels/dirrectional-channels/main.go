package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// NOTE: general to specific assigns
	cr = c
	cs = c

	// NOTE: general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// NOTE: specific to specific doesn't assign
	// cr = cs
	// cs = cr

	// NOTE: specific to specific doesn't convert
	// fmt.Println("-----")
	// fmt.Printf("c\t%T\n", (chan<- int)(cr))
	// fmt.Printf("c\t%T\n", (<-chan int)(cs))

	// NOTE: specific to general doesn't assign
	// c = cr
	// c = cs

	// NOTE: specific to general doesn't convert
	// fmt.Println("-----")
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan int)(cs))
}
