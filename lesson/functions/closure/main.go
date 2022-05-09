package main

import (
	"fmt"
)

// NOTE: A closure is a function that preserves the outer scope in its inner scope
// REF: https://www.javascripttutorial.net/javascript-closure/#:~:text=In%20JavaScript%2C%20a%20closure%20is,the%20lexical%20scoping%20works%20first.

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	// Below is Closure
	return func() int {
		x++
		return x
	}
}
