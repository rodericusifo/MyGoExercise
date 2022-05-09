package main

import (
	"fmt"
)

// Custom Type
type hotdog int

var a int
var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// Conversion
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
