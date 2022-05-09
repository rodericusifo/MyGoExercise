package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	// NOTE: (&) gives you the address
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	// NOTE: (*) gives you the value stored at an address when you have the address
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}
